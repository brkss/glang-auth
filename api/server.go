package api

import (
	"net/http"

	db "github.com/brkss/go-auth/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store db.Store
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{"response": "pong"})
	})

	server.router = router;
	return server
}

func (server *Server)Start(address string){
	server.router.Run(address)
}
