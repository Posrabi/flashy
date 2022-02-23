package entity

import "github.com/gocql/gocql"

type User struct {
	UserID       gocql.UUID `db:"user_id"`
	Username     string     `db:"user_name"`
	Name         string     `db:"name"`
	Email        string     `db:"email"`
	PhoneNumber  string     `db:"phone_number"`
	HashPassword string     `db:"hash_password"`
	AuthToken    string     `db:"auth_token"`
}
