package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/brkss/go-auth/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "payload"
)

// authMiddleware check authorization header get token and check its validity 
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc{
	return func(ctx *gin.Context){
		// get authorization header from request 
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := fmt.Errorf("authorization header not found !")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		// check authorization header fields 
		fields := strings.Split(authorizationHeader, " ")
		if len(fields) < 2 {
			err := fmt.Errorf("invalid authorization header !")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
		}

		// check token type from authorization header 
		if strings.ToLower(fields[0]) != authorizationTypeBearer {
			err := fmt.Errorf("invalid token type !")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return 
		}

		token := fields[1]
		payload, err := tokenMaker.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}
		// set payload in request context ! 
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
