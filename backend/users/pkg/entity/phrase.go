package entity

import (
	"time"

	proto "github.com/Posrabi/flashy/protos/users/proto"
)

type Phrase struct {
	UserID   string    `db:"user_id"`
	Word     string    `db:"word"`
	Sentence string    `db:"sentence"`
	Time     time.Time `db:"phrase_time"`
	Correct  bool      `db:"correct"`
}

func (p *Phrase) ToProto() *proto.Phrase {
	return &proto.Phrase{
		UserId:     p.UserID,
		Word:       p.Word,
		Sentence:   p.Sentence,
		PhraseTime: p.Time.Unix(),
		Correct:    p.Correct,
	}
}
