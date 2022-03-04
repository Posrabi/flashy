package middleware

import (
	"context"
	"os"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/golang-jwt/jwt/v4"

	"github.com/Posrabi/flashy/backend/common/pkg/auth"
)

type AuthLevel int

// nolint:revive,stylecheck
const (
	AuthLevel_NONE = iota
	AuthLevel_HIGH
)

var EndpointAuthMap = map[string]AuthLevel{
	"CreateUser": AuthLevel_NONE,
	"GetUser":    AuthLevel_HIGH,
	"UpdateUser": AuthLevel_HIGH,
	"DeleteUser": AuthLevel_NONE,
	"LogIn":      AuthLevel_NONE,
	"LogOut":     AuthLevel_NONE,
}

func NewJWTParser(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tokenString, ok := ctx.Value(kitjwt.JWTContextKey).(string)
			if !ok {
				_ = logger.Log("error", jwt.ErrTokenMalformed)
				return nil, jwt.ErrTokenMalformed
			}

			token, err := jwt.ParseWithClaims(tokenString, &auth.IDClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("ID_TOKEN")), nil
			})
			if err != nil {
				_ = logger.Log("error", jwt.ErrTokenInvalidClaims)
				return nil, jwt.ErrTokenInvalidClaims
			}
			return next(context.WithValue(ctx, kitjwt.JWTClaimsContextKey, token.Claims), req)
		}
	}
}
