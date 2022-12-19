package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.createCategoryParams{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	user, err := server.store.createCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) GetUser(ctx *gin.Context) {
	var req getUserRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	user, err := server.store.GetUser(ctx, req.Username)
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
