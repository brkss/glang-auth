package api

import (
	"net/http"

	db "github.com/brkss/go-auth/db/sqlc"
	token "github.com/brkss/go-auth/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store db.Store
	tokenMaker  token.Maker
}

func NewServer(store db.Store, tokenMaker token.Maker) *Server {
	server := &Server{store: store, tokenMaker: tokenMaker}
	router := gin.Default()
	
	router.POST("/login", server.Login)
	router.POST("/register", server.Register)

	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{"response": "pong"})
	})

	server.router = router;
	return server
}

func (server *Server)Start(address string){
	server.router.Run(address)
}

func errResponse(err error) gin.H{
	return gin.H{
		"error": err.Error(),
	}
}
