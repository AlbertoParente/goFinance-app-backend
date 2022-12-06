package api

import (
  "database/sql"
  "net/http"
  
  db "github.com/AlbertoParente/go-finance-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Email string `json:"email"`
}

func (server *Server) createUser(ctx *gin.Context) {
  var req createUserRequest
  err := ctx.ShouldBindJSON(&req)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, errorResponse(err))
  }
  
  arg := db.CreateUserParams{
    Username: req.Username,
    Password: req.Password,
    Email: req.Email,
  }
}
