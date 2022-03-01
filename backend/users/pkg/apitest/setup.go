package apitest

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"github.com/Posrabi/flashy/backend/common/pkg/utils"
)

func SetupEnv() {
	if _, exists := os.LookupEnv("SETUP"); exists {
		return
	}
	var ok bool
	UserEndpoint, ok = os.LookupEnv("USER_ENDPOINT")
	if !ok {
		UserEndpoint = "localhost:8080"
	}

	if err := godotenv.Load("../../../../build/.env"); err != nil && utils.FileExists("../../../../build/.env") {
		fmt.Println(".env loaded")
	}
	os.Setenv("SETUP", "COMPLETE")
}

func PopulateUser(t *testing.T) {
	for i, user := range TestUsers {
		userWithID, err := UserRepo.CreateUser(context.Background(), user)
		require.NoError(t, err)
		TestUsers[i] = userWithID
	}
}
