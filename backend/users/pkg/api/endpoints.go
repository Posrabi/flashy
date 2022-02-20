// Code generated .* DO NOT EDIT.
// To make changes, please modify codegen/endpoints.go.template

//nolint
package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
  "github.com/go-kit/kit/transport"

	proto "github.com/Posrabi/flashy/protos/users"
	logging "github.com/Posrabi/flashy/backend/middleware"
	commonservice "github.com/Posrabi/flashy/backend/common/pkg/service"
)

// Endpoints struct
type Endpoints struct {
	CreateUserEP endpoint.Endpoint
	GetUserEP endpoint.Endpoint
	UpdateUserEP endpoint.Endpoint
	DeleteUserEP endpoint.Endpoint
	AuthenticateEP endpoint.Endpoint
}

// CreateEndpoints creates endpoints
func CreateEndpoints(s Service, logger log.Logger) *Endpoints {
    handler := commonservice.NewLogErrorHandler(logger)

	var CreateUserEP endpoint.Endpoint
	{
		CreateUserEP = makeCreateUserEndpoint(s, handler)
    CreateUserEP = commonservice.AddRequestToContext("CreateUser")(CreateUserEP)
    CreateUserEP = logging.LoggingMiddleware(log.With(logger, "action", CreateUserEP))(CreateUserEP)
	}
	var GetUserEP endpoint.Endpoint
	{
		GetUserEP = makeGetUserEndpoint(s, handler)
    GetUserEP = commonservice.AddRequestToContext("GetUser")(GetUserEP)
    GetUserEP = logging.LoggingMiddleware(log.With(logger, "action", GetUserEP))(GetUserEP)
	}
	var UpdateUserEP endpoint.Endpoint
	{
		UpdateUserEP = makeUpdateUserEndpoint(s, handler)
    UpdateUserEP = commonservice.AddRequestToContext("UpdateUser")(UpdateUserEP)
    UpdateUserEP = logging.LoggingMiddleware(log.With(logger, "action", UpdateUserEP))(UpdateUserEP)
	}
	var DeleteUserEP endpoint.Endpoint
	{
		DeleteUserEP = makeDeleteUserEndpoint(s, handler)
    DeleteUserEP = commonservice.AddRequestToContext("DeleteUser")(DeleteUserEP)
    DeleteUserEP = logging.LoggingMiddleware(log.With(logger, "action", DeleteUserEP))(DeleteUserEP)
	}
	var AuthenticateEP endpoint.Endpoint
	{
		AuthenticateEP = makeAuthenticateEndpoint(s, handler)
    AuthenticateEP = commonservice.AddRequestToContext("Authenticate")(AuthenticateEP)
    AuthenticateEP = logging.LoggingMiddleware(log.With(logger, "action", AuthenticateEP))(AuthenticateEP)
	}
	return &Endpoints{
	    CreateUserEP: CreateUserEP,
	    GetUserEP: GetUserEP,
	    UpdateUserEP: UpdateUserEP,
	    DeleteUserEP: DeleteUserEP,
	    AuthenticateEP: AuthenticateEP,
	}
}

type createUserRequest struct {
	Request *proto.CreateUserRequest
}

type createUserResponse struct {
	Response *proto.CreateUserResponse
}

type getUserRequest struct {
	Request *proto.GetUserRequest
}

type getUserResponse struct {
	Response *proto.GetUserResponse
}

type updateUserRequest struct {
	Request *proto.UpdateUserRequest
}

type updateUserResponse struct {
	Response *proto.UpdateUserResponse
}

type deleteUserRequest struct {
	Request *proto.DeleteUserRequest
}

type deleteUserResponse struct {
	Response *proto.DeleteUserResponse
}

type authenticateRequest struct {
	Request *proto.AuthenticateRequest
}

type authenticateResponse struct {
	Response *proto.AuthenticateResponse
}


func makeCreateUserEndpoint(s Service, handler transport.ErrorHandler) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(createUserRequest)
		resp, err := s.CreateUser(ctx, req.Request)
    if err != nil {
      handler.Handle(ctx, err)
    }
		return createUserResponse{Response: resp}, err
	}
}

func makeGetUserEndpoint(s Service, handler transport.ErrorHandler) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(getUserRequest)
		resp, err := s.GetUser(ctx, req.Request)
    if err != nil {
      handler.Handle(ctx, err)
    }
		return getUserResponse{Response: resp}, err
	}
}

func makeUpdateUserEndpoint(s Service, handler transport.ErrorHandler) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(updateUserRequest)
		resp, err := s.UpdateUser(ctx, req.Request)
    if err != nil {
      handler.Handle(ctx, err)
    }
		return updateUserResponse{Response: resp}, err
	}
}

func makeDeleteUserEndpoint(s Service, handler transport.ErrorHandler) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(deleteUserRequest)
		resp, err := s.DeleteUser(ctx, req.Request)
    if err != nil {
      handler.Handle(ctx, err)
    }
		return deleteUserResponse{Response: resp}, err
	}
}

func makeAuthenticateEndpoint(s Service, handler transport.ErrorHandler) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(authenticateRequest)
		resp, err := s.Authenticate(ctx, req.Request)
    if err != nil {
      handler.Handle(ctx, err)
    }
		return authenticateResponse{Response: resp}, err
	}
}
