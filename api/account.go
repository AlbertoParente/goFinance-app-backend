package api

import (
	"database/sql"
	"net/http"

	db "github.com/AlbertoParente/go-finance-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.CreateAccountParams{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	user, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, user)
}

type GetAccountRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req GetAccountRequest
	err := ctx.ShouldBindUri(&req)

	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	user, err := server.store.GetAccount(ctx, req.Username)
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

type GetAccountByIdRequest struct {
	Username int32 `uri:"id" binding:"required"`
}

func (server *Server) GetAccountById(ctx *gin.Context) {
	var req GetAccountByIdRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	user, err := server.store.GetAccountById(ctx, req.ID)
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
