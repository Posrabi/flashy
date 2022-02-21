// Code generated .* DO NOT EDIT.
// To make changes, please modify codegen/transport.go.template

//nolint
package api

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
  "github.com/go-kit/kit/auth/jwt"
  kitlog "github.com/go-kit/kit/log"

	proto "github.com/Posrabi/flashy/protos/users"
)

type grpcServer struct {
	proto.UnimplementedUsersAPIServer
	CreateUserEP grpctransport.Handler
	GetUserEP grpctransport.Handler
	UpdateUserEP grpctransport.Handler
	DeleteUserEP grpctransport.Handler
	LogInEP grpctransport.Handler
	LogOutEP grpctransport.Handler
}

// NewGrpcTransport definition
func NewGrpcTransport(ep *Endpoints, logger kitlog.Logger) proto.UsersAPIServer {
	options := []grpctransport.ServerOption{
		//grpctransport.ServerErrorLogger(logger),
        grpctransport.ServerBefore(jwt.GRPCToContext()),
	}

	return &grpcServer{
		CreateUserEP: grpctransport.NewServer(
			ep.CreateUserEP,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
			options...,
		),
		GetUserEP: grpctransport.NewServer(
			ep.GetUserEP,
			decodeGetUserRequest,
			encodeGetUserResponse,
			options...,
		),
		UpdateUserEP: grpctransport.NewServer(
			ep.UpdateUserEP,
			decodeUpdateUserRequest,
			encodeUpdateUserResponse,
			options...,
		),
		DeleteUserEP: grpctransport.NewServer(
			ep.DeleteUserEP,
			decodeDeleteUserRequest,
			encodeDeleteUserResponse,
			options...,
		),
		LogInEP: grpctransport.NewServer(
			ep.LogInEP,
			decodeLogInRequest,
			encodeLogInResponse,
			options...,
		),
		LogOutEP: grpctransport.NewServer(
			ep.LogOutEP,
			decodeLogOutRequest,
			encodeLogOutResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	_, rep, err := s.CreateUserEP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.CreateUserResponse), nil
}

func decodeCreateUserRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.CreateUserRequest)
	return createUserRequest{Request: req}, nil
}

func encodeCreateUserResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(createUserResponse)
	return resp.Response, nil
}

func (s *grpcServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	_, rep, err := s.GetUserEP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.GetUserResponse), nil
}

func decodeGetUserRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.GetUserRequest)
	return getUserRequest{Request: req}, nil
}

func encodeGetUserResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(getUserResponse)
	return resp.Response, nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	_, rep, err := s.UpdateUserEP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.UpdateUserResponse), nil
}

func decodeUpdateUserRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.UpdateUserRequest)
	return updateUserRequest{Request: req}, nil
}

func encodeUpdateUserResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(updateUserResponse)
	return resp.Response, nil
}

func (s *grpcServer) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	_, rep, err := s.DeleteUserEP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.DeleteUserResponse), nil
}

func decodeDeleteUserRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.DeleteUserRequest)
	return deleteUserRequest{Request: req}, nil
}

func encodeDeleteUserResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(deleteUserResponse)
	return resp.Response, nil
}

func (s *grpcServer) LogIn(ctx context.Context, req *proto.LogInRequest) (*proto.LogInResponse, error) {
	_, rep, err := s.LogInEP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.LogInResponse), nil
}

func decodeLogInRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.LogInRequest)
	return logInRequest{Request: req}, nil
}

func encodeLogInResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(logInResponse)
	return resp.Response, nil
}

func (s *grpcServer) LogOut(ctx context.Context, req *proto.LogOutRequest) (*proto.LogOutResponse, error) {
	_, rep, err := s.LogOutEP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.LogOutResponse), nil
}

func decodeLogOutRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.LogOutRequest)
	return logOutRequest{Request: req}, nil
}

func encodeLogOutResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(logOutResponse)
	return resp.Response, nil
}
