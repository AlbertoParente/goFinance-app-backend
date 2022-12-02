package db

import (
	"context"
	"testing"

	"github.com/AlbertoParente/go-finance-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) account {
	user := createRandomUser(t)
	arq := CreateAccountParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	account, err := testQueries.CreateAccount(context.Background(), arq)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arq.UserID, account.UserID)
	require.Equal(t, arq.Title, account.Title)
	require.Equal(t, arq.Type, account.Type)
	require.Equal(t, arq.Description, account.Description)
	require.NotEmpty(t, user.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Title, account2.Title)
	require.Equal(t, account1.Description, account2.Description)
	require.NotEmpty(t, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}
