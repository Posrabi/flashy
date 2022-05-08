package api

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/go-kit/log"
	"google.golang.org/grpc/codes"

	gerr "github.com/Posrabi/flashy/backend/common/pkg/error"
	proto "github.com/Posrabi/flashy/protos/versus/proto"
)

const CHAN_BUF_SIZE = 5 // nolint: revive

// A global map that holds all of the channels. This should be initialized on server startup.
type queueMap struct {
	channels map[string]chan *proto.ConnectData // each client will get its own channel as a message queue
	mu       sync.Mutex
}

// Returns a new queue map.
func newQueueMap() *queueMap {
	return &queueMap{
		channels: map[string]chan *proto.ConnectData{},
	}
}

// Adds a new queue to the map, this is thread-safe.
func (q *queueMap) AddQueue(userID string) error {
	if !q.mu.TryLock() {
		return gerr.NewError(errors.New("unable to lock, this could case race conditions"), codes.Internal)
	}
	defer q.mu.Unlock()

	queue := make(chan *proto.ConnectData, CHAN_BUF_SIZE)
	_, ok := q.channels[userID]
	if ok {
		return gerr.NewError(errors.New("channels already exist"), codes.Internal)
	}

	q.channels[userID] = queue
	return nil
}

func (q *queueMap) GetQueue(userID string) (chan *proto.ConnectData, error) {
	if !q.mu.TryLock() {
		return nil, gerr.NewError(errors.New("unable to lock, this could case race conditions"), codes.Internal)
	}
	defer q.mu.Unlock()

	val, ok := q.channels[userID]
	if !ok {
		return nil, gerr.NewError(errors.New("channel not found"), codes.Internal)
	}

	return val, nil
}

// ConnectServer handles all of queueMap related operations.
type ConnectServer struct {
	logger log.Logger
	qm     *queueMap
}

// Returns a new ConnectServer.
func NewConnectServer(logger log.Logger) *ConnectServer {
	return &ConnectServer{
		logger: logger,
		qm:     newQueueMap(),
	}
}

// Join adds a new queue to the queueMap.
func (c *ConnectServer) Join(ctx context.Context, userID string) error {
	return c.qm.AddQueue(userID)
}

func (c *ConnectServer) Quit(ctx context.Context, userID string) error {
	q, err := c.qm.GetQueue(userID)
	if err != nil {
		return err
	}
	close(q)
	return nil
}

// Send takes msgs in the queue and send it to the client.
func (c *ConnectServer) Send(stream proto.VersusAPI_ConnectServer, userID string) error {
	channel, err := c.qm.GetQueue(userID)
	if err != nil {
		return err
	}

	for {
		msg, ok := <-channel
		if msg == nil && !ok { // channel has been closed and no remaining msgs.
			return nil
		}

		if err := stream.Send(msg); err != nil {
			return gerr.NewError(err, codes.Internal)
		}
	}
}

// Receive gather msgs sent by the client to the server and sent it to the according channel.
func (c *ConnectServer) Receive(stream proto.VersusAPI_ConnectServer) error {
	for {
		msg, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return gerr.NewError(err, codes.Internal)
		}

		queue, err := c.qm.GetQueue(msg.GetOpponentId())
		if err != nil {
			return err
		}
		queue <- msg
	}
}
