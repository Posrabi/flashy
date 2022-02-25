package scylla

import (
	"context"
	"errors"
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
	allColumns = "user_id, user_name, name, email, phone_number, hash_password"
)

func (u *userRepo) CreateUser(ctx context.Context, user *entity.User) (userID, authToken string, err error) {
	q := `INSERT INTO %s VALUES ($1, $2, $3, $4, $5, $6)`
	user.UserID, err = gocql.RandomUUID()
	if err != nil {
		return "", "", err
	}
	user.AuthToken, err = auth.GenerateToken(user.UserID)
	if err != nil {
		return "", "", err
	}
	return user.UserID.String(), user.AuthToken,
		u.sess.Query(fmt.Sprintf(q, tableName), user.UserID, user.Username, user.Name, user.Email,
			user.PhoneNumber, user.HashPassword).WithContext(ctx).Exec()
}

func (u *userRepo) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	q := `SELECT %s FROM %s WHERE user_id = $1`
	var user entity.User
	if err := u.sess.Query(fmt.Sprintf(q, allColumns, tableName), userID).WithContext(ctx).Idempotent(true).Scan(&user); err != nil {
		return nil, err
	}
	if user.AuthToken != ctx.Value(jwt.JWTContextKey) {
		return nil, errors.New("unauthorized request")
	}
	return &user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	q := `UPDATE %s SET (user_name = $1, name = $2, email = $3, phone_number = $4, hash_password = $5) 
	WHERE user_id = $6`
	if err := u.ValidateUser(ctx); err != nil {
		return err
	}
	return u.sess.Query(fmt.Sprintf(q, tableName), user.Username, user.Name, user.Email, user.PhoneNumber, user.HashPassword, user.UserID).
		WithContext(ctx).Idempotent(true).Exec()
}

func (u *userRepo) DeleteUser(ctx context.Context, user *entity.User) error {
	q := `DELETE FROM %s WHERE user_id = $1 AND hash_password = $2`
	return u.sess.Query(fmt.Sprintf(q, allColumns), user.UserID, user.HashPassword).WithContext(ctx).Consistency(gocql.All).Exec()
}

func (u *userRepo) LogIn(ctx context.Context, username, hashPassword string) (userID, authToken string, err error) {
	q := `SELECT user_id FROM %s WHERE user_name = $1 AND hash_password = $2`
	var uuid gocql.UUID
	if err := u.sess.Query(fmt.Sprintf(q, allColumns), username, hashPassword).Consistency(gocql.One).Scan(&uuid); err != nil {
		return uuid.String(), "", err
	}
	authToken, err = auth.GenerateToken(uuid)
	if err != nil {
		return uuid.String(), "", err
	}
	return uuid.String(), authToken, err
}

func (u *userRepo) LogOut(ctx context.Context, userID string) error {
	q := `DELETE auth_token FROM %s WHERE user_id = $1`
	return u.sess.Query(fmt.Sprintf(q, tableName), userID).Consistency(gocql.All).Exec()
}

func (u *userRepo) ValidateUser(ctx context.Context) error {
	q := `SELECT auth_token FROM %s WHERE user_id = $1`
	userID, err := auth.GetUserIDFromCtx(ctx)
	if err != nil {
		return err
	}
	var authToken string
	if err := u.sess.Query(fmt.Sprintf(q, tableName), userID).Idempotent(true).Scan(&authToken); err != nil {
		return err
	}
	if authToken != ctx.Value(jwt.JWTContextKey) {
		return errors.New("invalid token")
	}
	return nil
}
