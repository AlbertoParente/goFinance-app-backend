package api

import (
	"database/sql"
	"net/http"

	db "github.com/albertoparente/go-finance-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateAccountParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	}

	Account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, Account)
}

type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, erroResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type deleteAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	err = server.store.DeleteAccounts(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateAccountRequest struct {
	ID    int32  `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var req updateAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateAccountParams{
		ID:    req.ID,
		Title: req.Title,
	}

	account, err := server.store.UpdateAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountsRequest struct {
	UserID      int32  `json:"user_id" binding:"required`
	Type        string `json:"type" binding:"required`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountsRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, erroResponse(err))
	}

	arg := db.GetAccountsParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	}

	accounts, err := server.store.GetAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
