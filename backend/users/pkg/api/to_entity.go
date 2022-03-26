package api

import (
	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	proto "github.com/Posrabi/flashy/protos/users/proto"
)

func ConvertToUserEntity(user *proto.User) *entity.User {
	return &entity.User{
		UserID:       ConvertToUserIDEntity(user.GetUserId()),
		Username:     user.GetUserName(),
		Name:         user.GetName(),
		Email:        user.GetEmail(),
		HashPassword: user.GetHashPassword(),
		AuthToken:    user.GetAuthToken(),
	}
}

func ConvertToPhraseEntity(phrase *proto.Phrase) *entity.Phrase {
	return &entity.Phrase{
		UserID:   ConvertToUserIDEntity(phrase.UserId),
		Word:     phrase.GetWord(),
		Sentence: phrase.GetSentence(),
		Time:     phrase.GetPhraseTime().AsTime(),
	}
}

func ConvertToUserIDEntity(userID string) gocql.UUID {
	uuid, err := gocql.ParseUUID(userID)
	if err != nil {
		uuid = gocql.UUID{}
	}
	return uuid
}
