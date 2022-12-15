package api

import (
	"database/sql"
	"net/http"

	db "github.com/AlbertoParente/go-finance-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Category string `json:"category"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateCategoryParams{
		Category: req.Category,
	}

	user, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}

type GetCategoryRequest struct {
	Category string `uri:"category" binding:"required"`
}

func (server *Server) GetCategory(ctx *gin.Context) {
	var req GetCategoryRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	user, err := server.store.GetCategory(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, erroResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type GetCategoryByIdRequest struct {
	Category int32 `uri:"id" binding:"required"`
}

func (server *Server) GetCategoryById(ctx *gin.Context) {
	var req GetCategoryByIdRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	user, err := server.store.GetCategoryById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, erroResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
