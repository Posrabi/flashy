package apitest

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

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
