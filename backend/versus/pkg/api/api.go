package api

import (
	"context"
	"errors"

	"github.com/go-kit/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"

	gerr "github.com/Posrabi/flashy/backend/common/pkg/error"
	proto "github.com/Posrabi/flashy/protos/versus/proto"
)

type versusService struct {
	proto.UnimplementedVersusAPIServer
	cs     *connectServer
	logger log.Logger
}

func NewVersusService(logger log.Logger, qm *queueMap) proto.VersusAPIServer {
	return &versusService{
		cs:     newConnectServer(logger, qm),
		logger: logger,
	}
}

// Join joins a client to the server by creating a channel for it in queueMap.
func (vs *versusService) Join(ctx context.Context, req *proto.JoinRequest) (*proto.JoinResponse, error) {
	if err := vs.cs.Join(ctx, req.GetUserId()); err != nil {
		return nil, err
	}

	return &proto.JoinResponse{
		Success: true,
	}, nil
}

// Quit close the channel in queueMap.
func (vs *versusService) Quit(ctx context.Context, req *proto.QuitRequest) (*proto.QuitResponse, error) {
	if err := vs.cs.Quit(ctx, req.GetUserId()); err != nil {
		return nil, err
	}

	return &proto.QuitResponse{
		Success: true,
	}, nil
}

// Connect connects a client to the server, spawning two goroutines in the mean time.
func (vs *versusService) Connect(stream proto.VersusAPI_ConnectServer) error {
	group, _ := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return vs.cs.Receive(stream)
	})

	id := stream.Context().Value("user_id")
	userID, ok := id.(string)
	if !ok {
		return gerr.NewError(errors.New("unable to convert context id to string"), codes.InvalidArgument)
	}

	group.Go(func() error {
		return vs.cs.Send(stream, userID)
	})

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
