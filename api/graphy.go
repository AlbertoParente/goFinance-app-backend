//package api

//import (
//	"database/sql"
//	"net/http"

//	db "github.com/albertoparente/go-finance-app/db/sqlc"
//	"github.com/gin-gonic/gin"
//)

//type createGraphyCategoryRequest struct {
//	UserID      int32  `json:"user_id" binding:"required"`
//	Title       string `json:"title"`
//	Description string `json:"description"`
//}

//func (server *Server) createGraphyCategory(ctx *gin.Context) {
//	var req createGraphyCategoryRequest
//	err := ctx.ShouldBindJSON(&req)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//	}

//	arg := db.CreateGraphyCategoryParams{
//		UserID:      req.UserID,
//		Title:       req.Title,
//		Description: req.Description,
//	}

//	graphy, err := server.store.CreateGraphyCategory(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//	}

//	ctx.JSON(http.StatusOK, graphy)
//}
