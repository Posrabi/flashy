package apitest

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"github.com/Posrabi/flashy/backend/common/pkg/utils"
	"github.com/Posrabi/flashy/backend/users/pkg/api"
	"github.com/Posrabi/flashy/backend/users/pkg/scylla"
)

func Setup() {
	if _, exists := os.LookupEnv("SETUP"); exists {
		return
	}
	setupEnv()
	setupDB()
	os.Setenv("SETUP", "COMPLETE")
}

func setupEnv() {
	var ok bool
	UserEndpoint, ok = os.LookupEnv("USER_ENDPOINT")
	if !ok {
		UserEndpoint = "localhost:8080"
	}

	if err := godotenv.Load(); err != nil && utils.FileExists("../../../../build/.env") {
		fmt.Println(".env loaded")
	}
}

func setupDB() {
	sess, err := api.GetAccessToDB(api.ReadAndWrite, api.DevDB)
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	UserRepo = scylla.NewUserRepository(sess)
}

func PopulateUser(t *testing.T) {
	for _, user := range TestUsers {
		_, err := UserRepo.CreateUser(context.Background(), user)
		require.NoError(t, err)
		require.NoError(t, UpdateTestDataAfterCreate(CreateTestType_USER))
	}
}
