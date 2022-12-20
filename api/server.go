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

	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)
	router.GET("/category", server.createCategory)
	router.GET("/category/id/:id", server.getCategory)
	router.DELETE("/category/:id", server.deleteCategory)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func erroResponse(err error) gin.H {
	return gin.H("api has error:", err.Error())
}
