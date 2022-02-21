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

	"github.com/Posrabi/flashy/backend/common/pkg/utils"
	"github.com/Posrabi/flashy/backend/users/pkg/api"
	proto "github.com/Posrabi/flashy/protos/users"
)

var (
	grpcAddr = "localhost:8080"
)

func newServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "users gRPC server",
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
		fmt.Println("Exiting server on ", sig)
		os.Exit(0)
	}()

	fmt.Println("Starting gRPC server on", grpcAddr)
	if err := grpcServe(); err != nil {
		log.Panic(err)
	}
}

func grpcServe() error {
	// env := os.Getenv("ENV")

	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "timestamp", kitlog.DefaultTimestampUTC)

	var svcLogger = kitlog.With(logger, "component", "service")

	var svcUsers api.Service

	sess, err := api.GetAccessToDB(api.ReadAndWrite, api.UsersSpace)
	if err != nil {
		return err
	}
	defer sess.Close()

	svcUsers = api.NewService(api.NewMasterRepository(sess), svcLogger)
	svcUsers = api.NewLoggingService(kitlog.With(svcLogger, "service", "logger"), svcUsers)

	var (
		epsLogger   = kitlog.With(logger, "component", "endpoint")
		epsSvcUsers = api.CreateEndpoints(svcUsers, kitlog.With(epsLogger, "service", "users"))

		grpcLogger   = kitlog.With(logger, "component", "grpc")
		grpcSvcUsers = api.NewGrpcTransport(epsSvcUsers, grpcLogger)
	)

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}

	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Printf("Failed to close %s: %v\n", grpcAddr, err)
		}
	}()

	grpcServer := grpc.NewServer()
	proto.RegisterUsersAPIServer(grpcServer, grpcSvcUsers)

	return grpcServer.Serve(listener)
}
