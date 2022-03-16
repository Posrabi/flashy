package apitest

import (
	"time"

	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

var StartTime = time.Now().UTC()
var EndTime = time.Now().Add(10 * time.Minute).UTC() // nolint

var TestUsers = []*entity.User{
	{
		UserID:       gocql.UUID{},
		Username:     "test_user",
		Name:         "Test 1 2 3",
		Email:        "test@example.com",
		HashPassword: "thisisahash",
		AuthToken:    "supersecrettoken",
	},
	{
		UserID:       gocql.UUID{},
		Username:     "test_user_2",
		Name:         "1 2 3 Test",
		Email:        "test2@example.com",
		HashPassword: "thisisnotahash",
		AuthToken:    "supereasytoguesssecret",
	},
}
