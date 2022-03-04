package api

import (
	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	proto "github.com/Posrabi/flashy/protos/users/proto"
)

func ConvertToUserEntity(user *proto.User) *entity.User {
	uuid, err := gocql.ParseUUID(user.GetUserId())
	if err != nil {
		uuid = gocql.UUID{}
	}
	return &entity.User{
		UserID:       uuid,
		Username:     user.GetUserName(),
		Name:         user.GetName(),
		Email:        user.GetEmail(),
		PhoneNumber:  user.GetPhoneNumber(),
		HashPassword: user.GetHashPassword(),
		AuthToken:    user.GetAuthToken(),
	}
}
