package entity

import (
	"github.com/gocql/gocql"

	proto "github.com/Posrabi/flashy/protos/users/proto"
)

type User struct {
	UserID              gocql.UUID `db:"user_id"`
	Username            string     `db:"user_name"`
	Name                string     `db:"name"`
	Email               string     `db:"email"`
	HashPassword        string     `db:"hash_password"`
	FacebookAccessToken string     `db:"facebook_access_token"`
	AuthToken           string     `db:"auth_token"`
}

func (u *User) ToProto() *proto.User {
	return &proto.User{
		UserName:            u.Username,
		HashPassword:        u.HashPassword,
		Name:                u.Name,
		Email:               u.Email,
		AuthToken:           u.AuthToken,
		FacebookAccessToken: u.FacebookAccessToken,
		UserId:              u.UserID.String(),
	}
}
