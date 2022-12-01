package db

import (
	"context"
	"testing"

	"github.com/AlbertoParente/go-finance-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) account {
	user := createRandomUser(t)
	arq := CreateaccountParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	account, err := testQueries.Createaccount(context.Background(), arq)
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
	createRandomaccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomaccount(t)
	account2, err := testQueries.Getaccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Title, account2.Title)
	require.Equal(t, account1.Description, account2.Description)
	require.NotEmpty(t, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomaccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) account {
	account1 := createRandomaccount(t)

	arq := CreateaccountParams{
		ID:          account1.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(20),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arq)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arq.Title, account2.Title)
	require.Equal(t, arq.Description, account2.Description)
	require.NotEmpty(t, account2.CreatedAt)
}

func TestListAccount(t *testing.T) account {
	lasaccount := createRandomaccount(t)

	arq := CreateaccountParams{
		UserID:      lasaccount.UserID,
		Type:        lasaccount.Type,
		Title:       lasaccount.Title,
		Description: lasaccount.Description,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arq)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	for _, account := range accounts {
		require.Equal(t, lasaccount.ID, account2.ID)
		require.Equal(t, lasaccount.UserID, account2.UserID)
		require.Equal(t, lasaccount.Title, account2.Title)
		require.Equal(t, lasaccount.Description, account2.Description)
		require.NotEmpty(t, lasaccount.CreatedAt)
	}
}
