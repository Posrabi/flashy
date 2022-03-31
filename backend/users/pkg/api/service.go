package api

import (
	"context"
	"time"

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
		User: user.ToProto(),
	}, nil
}

func (s *service) GetUser(ctx context.Context, r *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := s.repo.GetUser(ctx, ConvertToUserIDEntity(r.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &proto.GetUserResponse{
		User: user.ToProto(),
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
	if err := s.repo.DeleteUser(ctx, ConvertToUserIDEntity(r.GetUserId())); err != nil {
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
		User: user.ToProto(),
	}, nil
}

func (s *service) LogOut(ctx context.Context, r *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	if err := s.repo.LogOut(ctx, ConvertToUserIDEntity(r.GetUserId())); err != nil {
		return nil, err
	}
	return &proto.LogOutResponse{
		Response: "Success",
	}, nil
}

func (s *service) CreatePhrase(ctx context.Context, r *proto.CreatePhraseRequest) (*proto.CreatePhraseResponse, error) {
	if err := s.repo.CreatePhrase(ctx, ConvertToPhraseEntity(r.GetPhrase())); err != nil {
		return nil, err
	}
	return &proto.CreatePhraseResponse{
		Response: "Success",
	}, nil
}

func (s *service) GetPhrases(ctx context.Context, r *proto.GetPhrasesRequest) (*proto.GetPhrasesResponse, error) {
	phrases, err := s.repo.GetPhrases(ctx, ConvertToUserIDEntity(r.GetUserId()), time.UnixMilli(r.GetStart()), time.UnixMilli(r.GetEnd()))
	if err != nil {
		return nil, err
	}
	var protoPhrases []*proto.Phrase
	for _, phrase := range phrases {
		protoPhrases = append(protoPhrases, phrase.ToProto())
	}

	return &proto.GetPhrasesResponse{
		Phrases: protoPhrases,
	}, nil
}

func (s *service) DeletePhrase(ctx context.Context, r *proto.DeletePhraseRequest) (*proto.DeletePhraseResponse, error) {
	if err := s.repo.DeletePhrase(ctx, ConvertToUserIDEntity(r.GetUserId()), time.UnixMilli(r.GetPhraseTime())); err != nil {
		return nil, err
	}

	return &proto.DeletePhraseResponse{
		Response: "Success",
	}, nil
}

func (s *service) LogInWithFB(ctx context.Context, r *proto.LogInWithFBRequest) (*proto.LogInWithFBResponse, error) {
	user, err := s.repo.LogInWithFB(ctx, ConvertToUserIDEntity(r.GetUserId()), r.GetFacebookAccessToken())
	if err != nil {
		return nil, err
	}

	return &proto.LogInWithFBResponse{
		User: user.ToProto(),
	}, nil
}
