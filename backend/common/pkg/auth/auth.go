package auth

import (
	"context"
	"errors"
	"os"
	"time"

	kitjwt "github.com/go-kit/kit/auth/jwt"
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
