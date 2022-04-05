package api

import (
	"time"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	proto "github.com/Posrabi/flashy/protos/users/proto"
)

func ConvertToUserEntity(user *proto.User) *entity.User {
	return &entity.User{
		UserID:              user.GetUserId(),
		Username:            user.GetUserName(),
		Name:                user.GetName(),
		Email:               user.GetEmail(),
		HashPassword:        user.GetHashPassword(),
		FacebookAccessToken: user.GetFacebookAccessToken(),
		AuthToken:           user.GetAuthToken(),
	}
}

func ConvertToPhraseEntity(phrase *proto.Phrase) *entity.Phrase {
	return &entity.Phrase{
		UserID:   phrase.UserId,
		Word:     phrase.GetWord(),
		Sentence: phrase.GetSentence(),
		Time:     time.UnixMilli(phrase.GetPhraseTime()),
		Correct:  phrase.GetCorrect(),
	}
}
