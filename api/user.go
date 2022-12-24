package api

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/brkss/go-auth/db/sqlc"
	"github.com/brkss/go-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type LoginRequest 		struct {
	Username	string `json:"username" binding:"required"`
	Password	string `json:"password" binding:"required"`
}

type RegisterRequest 	struct {
	Name		string `json:"name" binding:"required"`
	Username 	string `json:"username" binding:"required"`
	Password	string `json:"password" binding:"required"`
	Email		string `json:"email" binding:"required"`
} 

type AuthResponse struct {
	AccessToken	string `json:"access_token"`
}

func (server *Server)Login(ctx *gin.Context){

	var req LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	
}

func (server *Server)Register(ctx *gin.Context){

	var req RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	fmt.Println("calling register")
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateUserParams{
		ID: uuid.New().String(),
		Username: req.Username,
		Email: req.Email,
		Password: hash,
		Name: req.Name,
	}

	user, err := server.store.CreateUser(ctx, arg) 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	token, err := server.tokenMaker.CreateToken(user.ID, time.Hour)

	response := AuthResponse{
		AccessToken: token,
	}
	ctx.JSON(http.StatusOK, response)
	return;
}
