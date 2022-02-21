// Code generated .* DO NOT EDIT.
// To make changes, please modify codegen/logging.go.template

//nolint
package api

import (
	"context"

	"github.com/go-kit/log"
	proto "github.com/Posrabi/flashy/protos/users"
)

type loggingService struct {
	logging log.Logger
	Service
}

// NewLoggingService definition
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateUser(ctx context.Context, r *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return s.Service.CreateUser(ctx, r)
}

func (s *loggingService) GetUser(ctx context.Context, r *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	return s.Service.GetUser(ctx, r)
}

func (s *loggingService) UpdateUser(ctx context.Context, r *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	return s.Service.UpdateUser(ctx, r)
}

func (s *loggingService) DeleteUser(ctx context.Context, r *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	return s.Service.DeleteUser(ctx, r)
}

func (s *loggingService) LogIn(ctx context.Context, r *proto.LogInRequest) (*proto.LogInResponse, error) {
	return s.Service.LogIn(ctx, r)
}

func (s *loggingService) LogOut(ctx context.Context, r *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	return s.Service.LogOut(ctx, r)
}
