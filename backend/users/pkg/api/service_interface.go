// Code generated .* DO NOT EDIT.
// To make changes, please modify codegen/service_interface.go.template

//nolint
package api

import (
    "context"

    proto "github.com/Posrabi/flashy/protos/users"
)

type Service interface {
    CreateUser(ctx context.Context, r *proto.CreateUserRequest) (*proto.CreateUserResponse, error)
    GetUser(ctx context.Context, r *proto.GetUserRequest) (*proto.GetUserResponse, error)
    UpdateUser(ctx context.Context, r *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error)
    DeleteUser(ctx context.Context, r *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error)
    LogIn(ctx context.Context, r *proto.LogInRequest) (*proto.LogInResponse, error)
    LogOut(ctx context.Context, r *proto.LogOutRequest) (*proto.LogOutResponse, error)
}