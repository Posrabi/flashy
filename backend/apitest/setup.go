package apitest

import (
	"fmt"
	"os"
	"testing"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"

	"github.com/Posrabi/flashy/backend/common/pkg/utils"
	"github.com/Posrabi/flashy/backend/users/pkg/api"
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

func Setup(t *testing.T) *gocql.Session {
	SetupEnv()
	sess, err := api.SetupDB(api.ReadAndWrite, api.DevDB)
	if err != nil {
		panic(err)
	}
	go func() {
		defer sess.Close()
		<-sessCloseConn
	}()

	t.Cleanup(func() {
		sessCloseConn <- true
	})

	return sess
}
