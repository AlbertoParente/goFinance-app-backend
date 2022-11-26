package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arq := CreateUserParams{
		Username: "test",
		Password: "test123",
		Email:    "test@gmail.com",
	}

	user, err := testQueries.CreateUser(context.Background(), arq)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
