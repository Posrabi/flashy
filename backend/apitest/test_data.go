package apitest

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"

	"github.com/Posrabi/flashy/backend/users/pkg/api"
	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

var StartTime = time.Now().UTC()
var EndTime = time.Now().Add(10 * time.Minute).UTC() // nolint

var TestUsers = []*entity.User{
	{
		UserID:              gocql.UUID{},
		Username:            "test_user",
		Name:                "Test 1 2 3",
		Email:               "test@example.com",
		HashPassword:        "thisisahash",
		FacebookAccessToken: "v h7891y2189bfjds",
		AuthToken:           "supersecrettoken",
	},
	{
		UserID:              gocql.UUID{},
		Username:            "test_user_2",
		Name:                "1 2 3 Test",
		Email:               "test2@example.com",
		HashPassword:        "thisisnotahash",
		FacebookAccessToken: "423743298gggefdf",
		AuthToken:           "supereasytoguesssecret",
	},
}

var TestPhrases = []*entity.Phrase{
	{
		UserID:   api.ConvertToUserIDEntity(uuid.New().String()),
		Word:     "hello",
		Sentence: "hello world",
		Time:     time.Now(),
		Correct:  true,
	},
	{
		UserID:   api.ConvertToUserIDEntity(uuid.New().String()),
		Word:     "some word",
		Sentence: "some sentence",
		Time:     time.Now(),
		Correct:  true,
	},
}
