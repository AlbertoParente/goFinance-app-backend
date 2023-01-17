package api

import (
	"database/sql"
	"net/http"

	db "github.com/AlbertoParente/go-finance-app/db/sqlc"
	db "github.com/albertoparente/go-finance-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateCategoryParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
	}

	category, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, category)
}

type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}

	category, err := server.store.GetCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type deleteCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	var req deleteCategoryRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}

	err = server.store.DeleteCategories(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateCategoryRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) updateCategory(ctx *gin.Context) {
	var req updateCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateCategoryParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}

	category, err := server.store.UpdateCategories(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, category)
}

type getCategoriesRequest struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (server *Server) getCategories(ctx *gin.Context) {
	var req getCategoriesRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}

	var categories []db.Category
	var parametersHasUserIdAndType = req.UserID > 0 && len(req.Type) > 0

	filterAsByUserIdAnType := len(req.Description) == 0 && len(req.Title) == 0 && parametersHasUserIdAndType
	filterAsByUserIdAnType {
		arg := db.GetCategoriesByUserIdAndTypeParams{
			UserID: req.UserID,
			Type:   req.Type,
		}

		categoriesByUserIdAndType, err := server.store.GetCategoriesByUserIdAndType(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		categories = categoriesByUserIdAndType
	}

	filterAsByUserIdAndTypeAndDescription := len(req.Title) == 0 && len(req.Description) > 0 && parametersHasUserIdAndType
	if filterAsByUserIdAndTypeAndDescription {
		arg := db.GetCategoriesByUserIdAndTypeAnDescriptionParams{
			UserID:      req.UserID,
			Type:        req.Type,
			Description: req.Description,
		}

		categoriesByUserIdAndTypeAnDescription, err := server.store.GetCategoriesByUserIdAndTypeAnDescription(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		categories = categoriesByUserIdAndTypeAnDescription
	}

	filterAsByUserIdAndTypeAndTitle := len(req.Title) > 0 && len(req.Description) == 0 && parametersHasUserIdAndType
	if filterAsByUserIdAndTypeAndTitle {
		arg := db.GetCategoriesByUserIdAndTypeAndTitleParams{
			UserID: req.UserID,
			Type:   req.Type,
			Title:  req, Title,
		}

		categoriesByUserIdAndTypeAndTitle, err := server.store.GetCategoriesByUserIdAndTypeAndTitle(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		categories = categoriesByUserIdAndTypeAndTitle
	}
	
	filterAsAllParameters := len(req.Title) > 0 && len(req.Description) > 0 && parametersHasUserIdAndType
	if filterAsAllParameters {
		arg := db.GetCategoriesParams{
			UserID: req.UserID,
			Type:   req.Type,
			Title:  req, Title,
			Description: req.Description,
		}

		categoriesWithAllFilteres, err := server.store.GetCategories(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		categories = categoriesWithAllFilteres
	}

	ctx.JSON(http.StatusOK, categories)
}
