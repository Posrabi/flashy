package entity

import (
	"time"

	"github.com/gocql/gocql"

	proto "github.com/Posrabi/flashy/protos/users/proto"
)

type Phrase struct {
	UserID   gocql.UUID `db:"user_id"`
	Word     string     `db:"word"`
	Sentence string     `db:"sentence"`
	Time     time.Time  `db:"phrase_time"`
}

func (p *Phrase) ToProto() *proto.Phrase {
	return &proto.Phrase{
		UserId:     p.UserID.String(),
		Word:       p.Word,
		Sentence:   p.Sentence,
		PhraseTime: p.Time.Unix(),
	}
}
