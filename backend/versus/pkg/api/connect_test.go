package api

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []string{
	"1", "2",
}
var qMap = NewQueueMap()

func Test_AddQueue(t *testing.T) {
	wg := sync.WaitGroup{}
	for _, ch := range tests {
		wg.Add(1) // always add before the goroutine
		go func(id string) {
			defer wg.Done()

			require.NoError(t, newConnectServer(nil, qMap).qm.AddQueue(id))
		}(ch)
	}

	wg.Wait()
}

func Test_GetQueue(t *testing.T) {
	wg := sync.WaitGroup{}
	for _, ch := range tests {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			channel, err := newConnectServer(nil, qMap).qm.GetQueue(id)
			require.Nil(t, err)
			require.NotNil(t, channel)
		}(ch)
	}

	wg.Wait()
}

func Test_Quit(t *testing.T) {
	wg := sync.WaitGroup{}

	for _, ch := range tests {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			ctx := context.Background()
			cs := newConnectServer(nil, qMap)

			require.NoError(t, cs.Quit(ctx, id))

			c, err := cs.qm.GetQueue(id)
			require.Nil(t, err)
			require.Equal(t, 0, len(c))
		}(ch)
	}

	wg.Wait()
}
