package api

import (
	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	sc "github.com/Posrabi/flashy/backend/users/pkg/scylla"
)

type masterRepository struct {
	user repository.User
}

func NewMasterRepository(sess *gocql.Session) repository.Master {
	return &masterRepository{
		user: sc.NewUserRepository(sess),
	}
}
