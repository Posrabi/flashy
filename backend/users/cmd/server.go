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
	"github.com/Posrabi/flashy/backend/users/pkg/api"
	proto "github.com/Posrabi/flashy/protos/users/proto"
)

var addr = utils.GetNodeIPAddress() + port.USERS

// TODO: break up all of these functions.
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

	var svcLogger = kitlog.With(logger, "component", "service")

	var svcUsers api.Service

	dbType := api.DevDB
	if env := os.Getenv("ENV"); env == "prod" {
		dbType = api.ProdDB
	}

	sess, err := api.SetupDB(api.ReadAndWrite, dbType)
	if err != nil {
		return fmt.Errorf("failed to set up DB %w", err)
	}
	defer sess.Close()

	svcUsers = api.NewService(api.NewMasterRepository(sess), svcLogger)
	var (
		epsLogger   = kitlog.With(logger, "component", "endpoint")
		epsSvcUsers = api.CreateEndpoints(svcUsers, kitlog.With(epsLogger, "service", "users"))

		grpcLogger   = kitlog.With(logger, "component", "grpc")
		grpcSvcUsers = api.NewGrpcTransport(epsSvcUsers, grpcLogger)
	)

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return fmt.Errorf("failed to listen %w", err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Printf("Failed to close %s: %v\n", addr, err)
		}
	}()

	grpcServer := grpc.NewServer()
	proto.RegisterUsersAPIServer(grpcServer, grpcSvcUsers)

	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
