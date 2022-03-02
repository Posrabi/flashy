package auth

import (
	"context"
	"errors"
	"os"
	"time"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/gocql/gocql"
	"github.com/golang-jwt/jwt/v4"
)

type IDClaims struct {
	ID gocql.UUID `json:"id"`
	jwt.RegisteredClaims
}

const (
	flashy   = "Flashy"
	twoWeeks = 14
)

func GenerateToken(id gocql.UUID) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, IDClaims{
			ID:               id,
			RegisteredClaims: NewRegisteredClaims(),
		},
	)
	return token.SignedString([]byte(os.Getenv("ID_TOKEN")))
}

func NewRegisteredClaims() jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, twoWeeks)),
		Issuer:    flashy,
	}
}

func NewJWTParser() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tokenString, ok := ctx.Value(kitjwt.JWTContextKey).(string)
			if !ok {
				return nil, jwt.ErrTokenMalformed
			}
			token, err := jwt.ParseWithClaims(tokenString, &IDClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("ID_TOKEN")), nil
			})
			if err != nil {
				return nil, jwt.ErrTokenInvalidClaims
			}
			return next(context.WithValue(ctx, kitjwt.JWTClaimsContextKey, token.Claims), req)
		}
	}
}

func ValidateUserFromClaims(ctx context.Context, userID string) error {
	claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(*IDClaims)
	if !ok {
		return errors.New("missing claims")
	}

	if claims.ID.String() != userID {
		return errors.New("authentication error")
	}
	return nil
}

func ValidateUserFromToken(ctx context.Context, userID string) error {
	token, err := parseTokenFromContext(ctx)
	if err != nil {
		return err
	}
	return validateUser(token, userID)
}

func validateUser(token *jwt.Token, userID string) error {
	claims, ok := token.Claims.(*IDClaims)
	if !ok {
		return errors.New("missing claims")
	}

	if claims.ID.String() != userID {
		return errors.New("authentication error")
	}

	return nil
}
func parseTokenFromContext(ctx context.Context) (*jwt.Token, error) {
	tokenString, ok := ctx.Value(kitjwt.JWTContextKey).(string)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return jwt.ParseWithClaims(tokenString, &IDClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ID_TOKEN")), nil
	})
}
