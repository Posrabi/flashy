package auth

import (
	"context"
	"errors"
	"os"
	"time"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"

	gerr "github.com/Posrabi/flashy/backend/common/pkg/error"
)

type IDClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

const (
	flashy   = "Flashy"
	twoWeeks = 14
)

func GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		NewIDClaims(id),
	)
	return token.SignedString([]byte(os.Getenv("ID_TOKEN")))
}

func NewIDClaims(id string) *IDClaims {
	return &IDClaims{
		ID:               id,
		RegisteredClaims: NewRegisteredClaims(),
	}
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
		return gerr.NewError(errors.New("missing claims"), codes.Unauthenticated)
	}

	if claims.ID != userID {
		return gerr.NewError(errors.New("missing claims"), codes.PermissionDenied)
	}
	return nil
}
