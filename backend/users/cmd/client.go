package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/Posrabi/flashy/backend/common/pkg/utils"
	proto "github.com/Posrabi/flashy/protos/users"
)

func newClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "client services",
		Run:   runClientCmd,
	}
	return cmd
}

const address = "localhost:8081" // main svc

func runClientCmd(cmd *cobra.Command, args []string) {
	err := godotenv.Load()
	if err != nil && utils.FileExists(".env") {
		log.Print("Error loading .env file")
	}
	conn, err := grpc.Dial(address, grpc.WithAuthority(os.Getenv("INTERNAL_TOKEN")), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	_ = proto.NewUsersAPIClient(conn)

	_, cancel := context.WithTimeout(context.Background(), time.Second)
	cancel()
}
