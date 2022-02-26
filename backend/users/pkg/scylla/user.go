package scylla

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/common/pkg/auth"
	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	"github.com/Posrabi/flashy/backend/users/pkg/repository"
)

type userRepo struct {
	sess *gocql.Session
}

func NewUserRepository(sess *gocql.Session) repository.User {
	return &userRepo{
		sess: sess,
	}
}

// Consistency is Quorum by default.
const (
	tableName  = "users.info"
	allColumns = "user_id, user_name, name, email, phone_number, hash_password, auth_token"
)

// Create a brand new user, takes in a user without user_id and auth_token, returns a user with all values.
func (u *userRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	q := `INSERT INTO %s VALUES ($1, $2, $3, $4, $5, $6, $7)`
	userID, err := gocql.RandomUUID()
	if err != nil {
		return nil, err
	}
	authToken, err := auth.GenerateToken(user.UserID)
	if err != nil {
		return nil, err
	}
	if err := u.sess.Query(fmt.Sprintf(q, tableName), userID, user.Username, user.Name, user.Email,
		user.PhoneNumber, user.HashPassword, authToken).WithContext(ctx).Exec(); err != nil {
		return nil, err
	}
	user.UserID = userID
	user.AuthToken = authToken
	return user, nil
}

// GetUser to log in without user name and password. TODO: make this more secure.
func (u *userRepo) GetUser(ctx context.Context) (*entity.User, error) {
	q := `SELECT %s FROM %s WHERE auth_token = $1`
	var user entity.User
	if err := u.sess.Query(fmt.Sprintf(q, allColumns, tableName), ctx.Value(jwt.JWTContextKey)).WithContext(ctx).Idempotent(true).Scan(&user); err != nil {
		return nil, err
	}
	if err := auth.ValidateUser(ctx, user.UserID.String()); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	q := `UPDATE %s SET (user_name = $1, name = $2, email = $3, phone_number = $4, hash_password = $5) 
	WHERE user_id = $6`
	return u.sess.Query(fmt.Sprintf(q, tableName), user.Username, user.Name, user.Email, user.PhoneNumber, user.HashPassword, user.UserID).
		WithContext(ctx).Consistency(gocql.All).Idempotent(true).Exec()
}

func (u *userRepo) DeleteUser(ctx context.Context, userID, hashPassword string) error {
	q := `DELETE FROM %s WHERE user_id = $1 AND hash_password = $2`
	return u.sess.Query(fmt.Sprintf(q, allColumns), userID, hashPassword).WithContext(ctx).Consistency(gocql.One).Exec()
}

// LogIn with username and password.
func (u *userRepo) LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error) {
	q := `SELECT %s FROM %s WHERE user_name = $1 AND hash_password = $2`
	var user entity.User
	if err := u.sess.Query(fmt.Sprintf(q, allColumns, tableName), username, hashPassword).Consistency(gocql.One).Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// LogOut delete the auth token which is used to auto log in.
func (u *userRepo) LogOut(ctx context.Context, userID string) error {
	q := `DELETE auth_token FROM %s WHERE user_id = $1`
	return u.sess.Query(fmt.Sprintf(q, tableName), userID).Consistency(gocql.All).Exec()
}
