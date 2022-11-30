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
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	err := testQueries.DeleteCategories(context.Background(), category.ID)
	require.NoError(t, err)
}

func TestUpdateCategory(t *testing.T) Category {
	category1 := createRandomCategory(t)

	arq := CreateCategoryParams{
		ID:          category1.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(20),
	}

	category2, err := testQueries.UpdateCategories(context.Background(), arq)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arq.Title, category2.Title)
	require.Equal(t, arq.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestListCategories(t *testing.T) Category {
	lasCategory := createRandomCategory(t)

	arq := CreateCategoryParams{
		UserID:      lasCategory.UserID,
		Type:        lasCategory.Type,
		Title:       lasCategory.Title,
		Description: lasCategory.Description,
	}

	category2, err := testQueries.UpdateCategories(context.Background(), arq)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	for _, category := range categorys {
		require.Equal(t, lasCategory.ID, category2.ID)
		require.Equal(t, lasCategory.UserID, category2.UserID)
		require.Equal(t, lasCategory.Title, category2.Title)
		require.Equal(t, lasCategory.Description, category2.Description)
		require.NotEmpty(t, lasCategory.CreatedAt)
	}
}
