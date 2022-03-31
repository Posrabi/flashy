package scylla

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"

	"github.com/Posrabi/flashy/backend/common/pkg/auth"
	gerr "github.com/Posrabi/flashy/backend/common/pkg/error"
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
	info = "users.info"
)

// Create a brand new user, takes in a user without user_id and auth_token, returns a user with all values.
func (u *userRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	q := `INSERT INTO %s (user_id, user_name, name, email, hash_password, facebook_access_token, auth_token) VALUES (?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS`

	var err error
	user.UserID, err = gocql.RandomUUID()
	if err != nil {
		return nil, gerr.NewError(err, codes.Internal)
	}
	user.AuthToken, err = auth.GenerateToken(user.UserID)
	if err != nil {
		return nil, gerr.NewError(err, codes.Internal)
	}
	args := []interface{}{user.UserID, user.Username, user.Name, user.Email,
		user.HashPassword, user.FacebookAccessToken, user.AuthToken}

	if err := u.sess.Query(fmt.Sprintf(q, info), args...).Idempotent(true).WithContext(ctx).Exec(); err != nil {
		return nil, gerr.NewScError(err, codes.AlreadyExists, fmt.Sprintf(q, info), args)
	}

	return user, nil
}

// TODO: make this more secure.
func (u *userRepo) GetUser(ctx context.Context, userID gocql.UUID) (*entity.User, error) {
	q := `SELECT user_id, user_name, name, email, facebook_access_token, auth_token FROM %s WHERE user_id = ? AND auth_token = ?`

	var user entity.User
	if err := u.sess.Query(fmt.Sprintf(q, info), userID, ctx.Value(jwt.JWTContextKey)).WithContext(ctx).Consistency(gocql.One).Idempotent(true).Scan(
		&user.UserID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.FacebookAccessToken,
		&user.AuthToken,
	); err != nil {
		return nil, gerr.NewScError(err, codes.NotFound, fmt.Sprintf(q, info), []interface{}{userID.String()})
	}

	return &user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	q := `UPDATE %s SET user_name = ?, name = ?, email = ?, hash_password = ?, facebook_access_token = ? WHERE user_id = ?`

	if err := auth.ValidateUserFromClaims(ctx, user.UserID); err != nil {
		return err
	}

	if err := u.sess.Query(fmt.Sprintf(q, info), user.Username, user.Name, user.Email,
		user.HashPassword, user.FacebookAccessToken, user.UserID).Idempotent(true).WithContext(ctx).Exec(); err != nil {
		return gerr.NewScError(err, codes.Aborted, fmt.Sprintf(q, info), []interface{}{user.Username, user.Name, user.Email, user.HashPassword, user.FacebookAccessToken, user.UserID.String()})
	}
	return nil
}

func (u *userRepo) DeleteUser(ctx context.Context, userID gocql.UUID) error {
	q := `DELETE FROM %s WHERE user_id = ?`

	if err := u.sess.Query(fmt.Sprintf(q, info), userID).Exec(); err != nil {
		return gerr.NewScError(err, codes.NotFound, fmt.Sprintf(q, info), []interface{}{userID.String()})
	}
	return nil
}

func (u *userRepo) LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error) {
	q := `SELECT user_id, user_name, name, email, facebook_access_token, auth_token FROM %s WHERE user_name = ?`

	var user entity.User
	if err := u.sess.Query(fmt.Sprintf(q, info), username).Consistency(gocql.One).Idempotent(true).WithContext(ctx).Scan(
		&user.UserID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.FacebookAccessToken,
		&user.AuthToken,
	); err != nil {
		return nil, gerr.NewScError(err, codes.NotFound, fmt.Sprintf(q, info), []interface{}{
			username, hashPassword,
		})
	}
	if user.HashPassword != hashPassword {
		return nil, gerr.NewError(errors.New("invalid password"), codes.PermissionDenied)
	}

	return &user, nil
}

func (u *userRepo) LogOut(ctx context.Context, userID gocql.UUID) error {
	q := `UPDATE %s SET auth_token = ?, facebook_access_token = ? WHERE user_id = ? IF EXISTS`

	if err := u.sess.Query(fmt.Sprintf(q, info), " ", userID).Consistency(gocql.All).WithContext(ctx).Exec(); err != nil {
		return gerr.NewScError(err, codes.Aborted, fmt.Sprintf(q, info), []interface{}{" ", userID.String()})
	}
	return nil
}

func (u *userRepo) LogInWithFB(ctx context.Context, userID gocql.UUID, token string) (*entity.User, error) {
	q := `SELECT user_id, user_name, name, email, auth_token FROM %s WHERE user_id = ? AND facebook_access_token = ?`

	var user *entity.User
	if err := u.sess.Query(fmt.Sprintf(q, info), userID, token).WithContext(ctx).Consistency(gocql.One).Idempotent(true).Scan(
		&user.UserID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.AuthToken,
	); err != nil {
		return nil, gerr.NewScError(err, codes.NotFound, q, []interface{}{userID.String(), token})
	}

	return user, nil
}
