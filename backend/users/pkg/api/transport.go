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
	AuthenticateEP grpctransport.Handler
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
		AuthenticateEP: grpctransport.NewServer(
			ep.AuthenticateEP,
			decodeAuthenticateRequest,
			encodeAuthenticateResponse,
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

func (s *grpcServer) Authenticate(ctx context.Context, req *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	_, rep, err := s.AuthenticateEP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.AuthenticateResponse), nil
}

func decodeAuthenticateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*proto.AuthenticateRequest)
	return authenticateRequest{Request: req}, nil
}

func encodeAuthenticateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(authenticateResponse)
	return resp.Response, nil
}
