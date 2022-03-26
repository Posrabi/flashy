package api

import (
	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	proto "github.com/Posrabi/flashy/protos/users/proto"
)

func ConvertToUserProto(user *entity.User) *proto.User {
	return &proto.User{
		UserName:     user.Username,
		HashPassword: user.HashPassword,
		Name:         user.Name,
		Email:        user.Email,
		AuthToken:    user.AuthToken,
		UserId:       user.UserID.String(),
	}
}

func ConvertToPhraseProto(phrase *entity.Phrase) *proto.Phrase {
	return &proto.Phrase{
		UserId:     phrase.UserID.String(),
		Word:       phrase.Word,
		Sentence:   phrase.Sentence,
		PhraseTime: phrase.Time.Unix(),
	}
}
