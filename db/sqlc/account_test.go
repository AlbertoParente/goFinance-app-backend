package db

import (
	"context"
	"testing"
 	"time"

 	"github.com/AlbertoParente/go-finance-app/util"
 	"github.com/stretchr/testify/require"
)

 func createRandomAccount(t *testing.T) Account {
 	category := createRandomCategory(t)
 	arg := CreateAccountParams{
		UserID:      category.UserID,
 		CategoryID:  category.ID,
 		Title:       util.RandomString(12),
 		Type:        category.Type,
		Description: util.RandomString(20),
		Value:       10,
		Date:        time.Now(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.CategoryID, account.CategoryID)
	require.Equal(t, arg.Value, account.Value)
	require.Equal(t, arg.Title, account.Title)
	require.Equal(t, arg.Type, account.Type)
	require.Equal(t, arg.Description, account.Description)
	require.NotEmpty(t, account.CreatedAt)
 	require.NotEmpty(t, account.Date)

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

	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.CategoryID, account2.CategoryID)
 	require.Equal(t, account1.Value, account2.Value)
 	require.Equal(t, account1.Title, account2.Title)
 	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Description, account2.Description)
 	require.NotEmpty(t, account2.Date)
	require.NotEmpty(t, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
 	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
 }

// func TestListGetReports(t *testing.T) {
// 	lastAccount := createRandomAccount(t)

// 	arg := GetAccountsReportsParams{
// 		UserID: lastAccount.UserID,
// 		Type:   lastAccount.Type,
// 	}

// 	sumValue, err := testQueries.GetAccountsReports(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, sumValue)
// }

// func TestListGetGraph(t *testing.T) {
// 	lastAccount := createRandomAccount(t)

// 	arg := GetAccountsGraphParams{
// 		UserID: lastAccount.UserID,
// 		Type:   lastAccount.Type,
// 	}

// 	graphValue, err := testQueries.GetAccountsGraph(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, graphValue)
// }
