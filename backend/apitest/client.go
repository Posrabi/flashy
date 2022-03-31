package apitest

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "github.com/Posrabi/flashy/protos/users/proto"
)

var (
	sessCloseConn = make(chan bool, 1)
	closeConn     = make(chan int, 1)
	UserAPI       proto.UsersAPIClient
)

func SetupAPIConnection(t *testing.T) {
	conn := setupAPIConn(t)
	cleanUpAfterTest(t, conn)
}

func setupAPIConn(t *testing.T) *grpc.ClientConn {
	conn, err := grpc.Dial(UserEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Logf("this test requires a live gRPC connection to the api - error: %s", err.Error())
		t.Fail()
	}
	UserAPI = proto.NewUsersAPIClient(conn)
	return conn
}

func cleanUpAfterTest(t *testing.T, conn *grpc.ClientConn) {
	go func() {
		defer conn.Close()
		<-closeConn
	}()

	t.Cleanup(func() {
		closeConn <- 1
	})
}
