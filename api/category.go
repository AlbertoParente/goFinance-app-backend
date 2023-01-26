package api

import (
	"database/sql"
	"net/http"

	db "github.com/AlbertoParente/go-finance-app/db/sqlc"
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
