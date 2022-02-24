package auth

import (
	"context"
	"fmt"
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

var IDSigningMethod jwt.SigningMethodHMAC

const (
	flashy   = "Flashy"
	twoWeeks = 14
)

func (c *IDClaims) Valid() error {
	if c.ID.String() == "" {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func GenerateToken(id gocql.UUID) (string, error) {
	token := jwt.NewWithClaims(
		&IDSigningMethod, &IDClaims{
			ID:               id,
			RegisteredClaims: newRegisteredClaims(),
		},
	)
	return token.SignedString(os.Getenv("ID_TOKEN"))
}

func newRegisteredClaims() jwt.RegisteredClaims {
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
