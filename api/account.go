package api

import (
	"database/sql"
	"net/http"

	db "github.com/AlbertoParente/go-finance-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Username string `json:"username"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.CreateAccountParams{
		Username: req.Username,
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
