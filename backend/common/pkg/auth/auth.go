package auth

import (
	"fmt"
	"os"

	"github.com/gocql/gocql"
	"github.com/golang-jwt/jwt/v4"
)

type IDClaims struct {
	ID gocql.UUID `json:"id"`
	jwt.RegisteredClaims
}

var IDSigningMethod = *jwt.SigningMethodES256

func (c *IDClaims) Valid() error {
	if c.ID.String() == "" {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func GenerateToken(id gocql.UUID) (string, error) {
	token := jwt.NewWithClaims(
		&IDSigningMethod, &IDClaims{
			ID: id,
		},
	)
	return token.SignedString([]byte(os.Getenv("ID_TOKEN")))
}
