package api

import (
	"context"

	"github.com/go-kit/log"

	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	proto "github.com/Posrabi/flashy/protos/users/proto"
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
	user, err := s.repo.CreateUser(ctx, ConvertToUserEntity(r.GetUser()))
	if err != nil {
		return nil, err
	}
	return &proto.CreateUserResponse{
		User: ConvertToUserProto(user),
	}, nil
}

func (s *service) GetUser(ctx context.Context, r *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := s.repo.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return &proto.GetUserResponse{
		User: ConvertToUserProto(user),
	}, nil
}

func (s *service) UpdateUser(ctx context.Context, r *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	if err := s.repo.UpdateUser(ctx, ConvertToUserEntity(r.GetUser())); err != nil {
		return nil, err
	}
	return &proto.UpdateUserResponse{
		Response: "Success",
	}, nil
}

func (s *service) DeleteUser(ctx context.Context, r *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	if err := s.repo.DeleteUser(ctx, r.GetUserId(), r.GetHashPassword()); err != nil {
		return nil, err
	}
	return &proto.DeleteUserResponse{
		Response: "Success",
	}, nil
}

func (s *service) LogIn(ctx context.Context, r *proto.LogInRequest) (*proto.LogInResponse, error) {
	user, err := s.repo.LogIn(ctx, r.GetUserName(), r.GetHashPassword())
	if err != nil {
		return nil, err
	}
	return &proto.LogInResponse{
		User: ConvertToUserProto(user),
	}, nil
}

func (s *service) LogOut(ctx context.Context, r *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	if err := s.repo.LogOut(ctx, r.GetUserId()); err != nil {
		return nil, err
	}
	return &proto.LogOutResponse{
		Response: "Success",
	}, nil
}
