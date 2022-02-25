package api

import (
	"context"

	"github.com/go-kit/log"

	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	proto "github.com/Posrabi/flashy/protos/users"
)

type service struct {
	repo   repository.Master
	logger log.Logger
}

func NewService(r repository.Master, logger log.Logger) Service {
	return &service{
		repo:   r,
		logger: logger,
	}
}

func (s *service) CreateUser(ctx context.Context, r *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return nil, nil
}

func (s *service) GetUser(ctx context.Context, r *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	return nil, nil
}

func (s *service) UpdateUser(ctx context.Context, r *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	return nil, nil
}

func (s *service) DeleteUser(ctx context.Context, r *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	return nil, nil
}

func (s *service) LogIn(ctx context.Context, r *proto.LogInRequest) (*proto.LogInResponse, error) {
	return nil, nil
}

func (s *service) LogOut(ctx context.Context, r *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	return nil, nil
}
