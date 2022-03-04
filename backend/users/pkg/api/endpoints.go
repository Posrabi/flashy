// Code generated .* DO NOT EDIT.
// To make changes, please modify codegen/endpoints.go.template

//nolint
package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
  "github.com/go-kit/kit/transport"

	proto "github.com/Posrabi/flashy/protos/users/proto"
	logging "github.com/Posrabi/flashy/backend/middleware"
	"github.com/Posrabi/flashy/backend/middleware"
	commonauth "github.com/Posrabi/flashy/backend/common/pkg/auth"
)

// Endpoints struct
type Endpoints struct {
	CreateUserEP endpoint.Endpoint
	GetUserEP endpoint.Endpoint
	UpdateUserEP endpoint.Endpoint
	DeleteUserEP endpoint.Endpoint
	LogInEP endpoint.Endpoint
	LogOutEP endpoint.Endpoint
}

// CreateEndpoints creates endpoints
func CreateEndpoints(s Service, logger log.Logger) *Endpoints {
    handler := middleware.NewLogErrorHandler(logger)

	var CreateUserEP endpoint.Endpoint
	{
		CreateUserEP = makeCreateUserEndpoint(s, handler)
    CreateUserEP = middleware.AddRequestToContext("CreateUser")(CreateUserEP)
		CreateUserEP = commonauth.NewJWTParser()(CreateUserEP)
    CreateUserEP = logging.LoggingMiddleware(log.With(logger, "action", CreateUserEP))(CreateUserEP)

	}
	var GetUserEP endpoint.Endpoint
	{
		GetUserEP = makeGetUserEndpoint(s, handler)
    GetUserEP = middleware.AddRequestToContext("GetUser")(GetUserEP)
		GetUserEP = commonauth.NewJWTParser()(GetUserEP)
    GetUserEP = logging.LoggingMiddleware(log.With(logger, "action", GetUserEP))(GetUserEP)

	}
	var UpdateUserEP endpoint.Endpoint
	{
		UpdateUserEP = makeUpdateUserEndpoint(s, handler)
    UpdateUserEP = middleware.AddRequestToContext("UpdateUser")(UpdateUserEP)
		UpdateUserEP = commonauth.NewJWTParser()(UpdateUserEP)
    UpdateUserEP = logging.LoggingMiddleware(log.With(logger, "action", UpdateUserEP))(UpdateUserEP)

	}
	var DeleteUserEP endpoint.Endpoint
	{
		DeleteUserEP = makeDeleteUserEndpoint(s, handler)
    DeleteUserEP = middleware.AddRequestToContext("DeleteUser")(DeleteUserEP)
		DeleteUserEP = commonauth.NewJWTParser()(DeleteUserEP)
    DeleteUserEP = logging.LoggingMiddleware(log.With(logger, "action", DeleteUserEP))(DeleteUserEP)

	}
	var LogInEP endpoint.Endpoint
	{
		LogInEP = makeLogInEndpoint(s, handler)
    LogInEP = middleware.AddRequestToContext("LogIn")(LogInEP)
		LogInEP = commonauth.NewJWTParser()(LogInEP)
    LogInEP = logging.LoggingMiddleware(log.With(logger, "action", LogInEP))(LogInEP)

	}
	var LogOutEP endpoint.Endpoint
	{
		LogOutEP = makeLogOutEndpoint(s, handler)
    LogOutEP = middleware.AddRequestToContext("LogOut")(LogOutEP)
		LogOutEP = commonauth.NewJWTParser()(LogOutEP)
    LogOutEP = logging.LoggingMiddleware(log.With(logger, "action", LogOutEP))(LogOutEP)

	}
	return &Endpoints{
	    CreateUserEP: CreateUserEP,
	    GetUserEP: GetUserEP,
	    UpdateUserEP: UpdateUserEP,
	    DeleteUserEP: DeleteUserEP,
	    LogInEP: LogInEP,
	    LogOutEP: LogOutEP,
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

type logInRequest struct {
	Request *proto.LogInRequest
}

type logInResponse struct {
	Response *proto.LogInResponse
}

type logOutRequest struct {
	Request *proto.LogOutRequest
}

type logOutResponse struct {
	Response *proto.LogOutResponse
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

func makeLogInEndpoint(s Service, handler transport.ErrorHandler) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(logInRequest)
		resp, err := s.LogIn(ctx, req.Request)
    if err != nil {
      handler.Handle(ctx, err)
    }
		return logInResponse{Response: resp}, err
	}
}

func makeLogOutEndpoint(s Service, handler transport.ErrorHandler) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(logOutRequest)
		resp, err := s.LogOut(ctx, req.Request)
    if err != nil {
      handler.Handle(ctx, err)
    }
		return logOutResponse{Response: resp}, err
	}
}
