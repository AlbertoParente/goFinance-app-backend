package db

import (
	"context"
	"testing"

	"github.com/AlbertoParente/go-finance-app/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arq := CreateCategoryParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arq)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arq.UserID, category.UserID)
	require.Equal(t, arq.Title, category.Title)
	require.Equal(t, arq.Type, category.Type)
	require.Equal(t, arq.Description, category.Description)
	require.NotEmpty(t, user.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	user1 := createRandomCategory(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}
