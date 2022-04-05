package scylla_test

import (
	"context"
	"testing"

	"github.com/Posrabi/flashy/backend/apitest"
	"github.com/Posrabi/flashy/backend/common/pkg/auth"
	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	"github.com/Posrabi/flashy/backend/users/pkg/scylla"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	userSetup(t)

	t.Run("Create_User", func(t *testing.T) {
		testCreate_User(t, apitest.UserRepo)
	})

	t.Run("LogIn_User", func(t *testing.T) {
		testLogIn_User(t, apitest.UserRepo)
	})

	t.Run("LogInWithFB_User", func(t *testing.T) {
		testLogInWithFB_User(t, apitest.UserRepo)
	})

	t.Run("Get_User", func(t *testing.T) {
		testGet_User(t, apitest.UserRepo)
	})

	t.Run("Update_User", func(t *testing.T) {
		testUpdate_User(t, apitest.UserRepo)
	})

	t.Run("LogOut_User", func(t *testing.T) {
		testLogOut_User(t, apitest.UserRepo)
	})

	t.Run("Delete_User", func(t *testing.T) {
		testDelete_User(t, apitest.UserRepo)
	})
}

func userSetup(t *testing.T) {
	sess := apitest.Setup(t)

	apitest.UserRepo = scylla.NewUserRepository(sess)

	t.Cleanup(func() {
		for _, user := range apitest.TestUsers {
			require.NoError(t, apitest.UserRepo.DeleteUser(context.Background(), user.UserID))
		}
	})

}

func testCreate_User(t *testing.T, repo repository.User) {
	t.Helper()

	for i, user := range apitest.TestUsers {
		updatedUser, err := repo.CreateUser(context.Background(), user)
		require.NoError(t, err)
		apitest.TestUsers[i] = updatedUser
	}
}

func testGet_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, expected := range apitest.TestUsers {
		actual, err := repo.GetUser(context.WithValue(context.Background(), jwt.JWTContextKey, expected.AuthToken), expected.UserID)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}
}

func testUpdate_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, expected := range apitest.TestUsers {
		expected.Email = "newemail@example.com"
		expected.Name = "update user tester"
		expected.Username = "new_user"
		expected.HashPassword = "newpassword"
		expected.FacebookAccessToken = "newToken"

		ctx := context.WithValue(context.Background(), jwt.JWTClaimsContextKey, auth.NewIDClaims(expected.UserID))
		require.NoError(t, repo.UpdateUser(ctx, expected))

		actual, err := repo.GetUser(ctx, expected.UserID)
		require.NoError(t, err)
		expected.HashPassword = actual.HashPassword
		require.Equal(t, expected, actual)
	}
}

func testDelete_User(t *testing.T, repo repository.User) {
	t.Helper()

	for i, deleting := range apitest.TestUsers {
		require.NoError(t, repo.DeleteUser(context.Background(), deleting.UserID))
		removeUserAtIndex(apitest.TestUsers, i)
	}
}

func testLogIn_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, expected := range apitest.TestUsers {
		actual, err := repo.LogIn(context.Background(), expected.Username, expected.HashPassword)
		require.NoError(t, err)
		require.Equal(t, expected.UserID, actual.UserID)
		require.Equal(t, expected.AuthToken, actual.AuthToken)
	}
}

func testLogOut_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, user := range apitest.TestUsers {
		require.NoError(t, repo.LogOut(context.WithValue(context.Background(), jwt.JWTContextKey,
			user.AuthToken), user.UserID))
	}
}

func testLogInWithFB_User(t *testing.T, repo repository.User) {
	t.Helper()

	for _, expected := range apitest.TestUsers {
		if expected.FacebookAccessToken != "" {
			actual, err := repo.LogInWithFB(context.Background(), expected.UserID, expected.FacebookAccessToken)
			require.NoError(t, err)
			actual.Username = expected.Username
			actual.HashPassword = expected.HashPassword
			require.Equal(t, expected, actual)
		}
	}
}

func removeUserAtIndex(users []*entity.User, i int) []*entity.User {
	return append(users[:i], users[i+1:]...)
}
