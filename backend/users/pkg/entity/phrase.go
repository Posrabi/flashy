package entity

import (
	"time"

	"github.com/gocql/gocql"
)

type Phrase struct {
	UserID   gocql.UUID `db:"user_id"`
	Word     string     `db:"word"`
	Sentence string     `db:"sentence"`
	Time     time.Time  `db:"cur_time"`
}
