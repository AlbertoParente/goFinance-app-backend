package api

import (
	db "github.com/AlbertoParente/go-finance-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func newServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func erroResponse(err error) gin.H {
	return gin.H("api has error:", err.Error())
}
