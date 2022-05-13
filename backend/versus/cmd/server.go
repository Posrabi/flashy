package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/log"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	port "github.com/Posrabi/flashy/backend/common/pkg/ports"
	"github.com/Posrabi/flashy/backend/common/pkg/utils"
	"github.com/Posrabi/flashy/backend/versus/pkg/api"
	proto "github.com/Posrabi/flashy/protos/versus/proto"
)

var addr = utils.GetNodeIPAddress() + port.VERSUS

func newServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "versus gRPC server",
		Run:   runServerCmd,
	}
	return cmd
}

func runServerCmd(cmd *cobra.Command, args []string) {
	err := godotenv.Load()
	if err != nil && utils.FileExists(".env") {
		log.Print("Error loading .env file")
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		cancel()
		fmt.Println(" Exiting server on ", sig)
		os.Exit(0)
	}()

	fmt.Println("Starting gRPC server on", addr)
	if err := grpcServe(); err != nil {
		log.Panic(err)
	}
}

func grpcServe() error {
	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "timestamp", kitlog.DefaultTimestampUTC)

	var svcLogger = kitlog.With(logger, "componenet", "service")

	var svcVersus = api.NewVersusService(svcLogger, api.NewQueueMap())

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen %w", err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Printf("Failed to close %s: %v\n", addr, err)
		}
	}()

	grpcServer := grpc.NewServer()
	proto.RegisterVersusAPIServer(grpcServer, svcVersus)

	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
