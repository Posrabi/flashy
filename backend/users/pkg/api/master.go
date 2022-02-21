package api

import (
	"github.com/scylladb/gocqlx/v2"

	"github.com/Posrabi/flashy/backend/users/pkg/repository"
)

type masterRepository struct {
}

func NewMasterRepository(gocqlx.Session) repository.Master {
	return &masterRepository{}
}
