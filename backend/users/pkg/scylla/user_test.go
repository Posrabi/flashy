package scylla_test

import (
	"context"
	"testing"

	"github.com/Posrabi/flashy/backend/users/pkg/apitest"
	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	userSetup(t)

	t.Run("Create_User", func(t *testing.T) {
		testCreate_User(t, apitest.UserRepo)
	})

	t.Run("Get_User", func(t *testing.T) {
		testGet_User(t, apitest.UserRepo)
	})

	t.Run("LogIn_User", func(t *testing.T) {
		testLogIn_User(t, apitest.UserRepo)
	})

	t.Run("LogOut_User", func(t *testing.T) {
		testLogOut_User(t, apitest.UserRepo)
	})

	t.Run("Delete_User", func(t *testing.T) {
		testDelete_User(t, apitest.UserRepo)
	})
}

func userSetup(t *testing.T) {
	apitest.Setup()

	apitest.PopulateUser(t)

	t.Cleanup(func() {
		for _, user := range apitest.TestUsers {
			require.NoError(t, apitest.UserRepo.DeleteUser(context.Background(),
				user.UserID.String(), user.HashPassword))
		}
	})
}

func testCreate_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, user := range apitest.TestUsers {
		_, err := repo.CreateUser(context.Background(), user)
		require.NoError(t, err)
	}
}

func testGet_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, expected := range apitest.TestUsers {
		ctx := context.WithValue(context.Background(), jwt.JWTContextKey, expected.AuthToken)
		actual, err := repo.GetUser(ctx)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}
}

func testUpdate_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, expected := range apitest.TestUsers {
		newUserId, err := gocql.ParseUUID(uuid.New().String())
		require.NoError(t, err)
		expected.UserID = newUserId
		expected.Email = "newemail@example.com"
		expected.AuthToken = "haha"
		expected.Name = "update user tester"
		expected.Username = "new_user"
		expected.PhoneNumber = "+16476666666"
		require.NoError(t, repo.UpdateUser(context.Background(), expected))

		actual, err := repo.LogIn(context.Background(), expected.Username, expected.HashPassword)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}
}

func testDelete_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, deleting := range apitest.TestUsers {
		require.NoError(t, repo.DeleteUser(context.Background(),
			deleting.UserID.String(), deleting.HashPassword))

		actualUser, err := repo.LogIn(context.Background(), deleting.Username, deleting.HashPassword)
		require.NoError(t, err)
		require.Equal(t, nil, actualUser)
	}
}

func testLogIn_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, expected := range apitest.TestUsers {
		actual, err := repo.LogIn(context.Background(), expected.Username, expected.HashPassword)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}
}

func testLogOut_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, user := range apitest.TestUsers {
		require.NoError(t, repo.LogOut(context.Background(), user.UserID.String()))
		ctx := context.WithValue(context.Background(), jwt.JWTContextKey, user.AuthToken)
		actual, err := repo.GetUser(ctx)
		require.NoError(t, err)
		require.Equal(t, nil, actual)
	}
}
