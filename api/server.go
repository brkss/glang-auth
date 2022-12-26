package api

import (
	db "github.com/brkss/go-auth/db/sqlc"
	token "github.com/brkss/go-auth/token"
	"github.com/brkss/go-auth/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	store      db.Store
	tokenMaker token.Maker
	config     *utils.Config
}

func NewServer(store db.Store, tokenMaker token.Maker, config *utils.Config) *Server {
	server := &Server{store: store, tokenMaker: tokenMaker, config: config}
	router := gin.Default()

	router.POST("/login", server.Login)
	router.POST("/register", server.Register)

	// -- Protected routes
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/me", server.Me)

	server.router = router
	return server
}

func (server *Server) Start(address string) {
	server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
