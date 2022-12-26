package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brkss/go-auth/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)




func TestAuthMiddleware(t *testing.T){

	testCases := []struct{
		name string
		setAuth func(request *http.Request, tokenMaker token.Maker)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setAuth: func(request *http.Request, tokenMaker token.Maker){
				token, err := tokenMaker.CreateToken(uuid.New().String(), time.Minute)					
				require.NoError(t, err)

				authorization := fmt.Sprintf("%s %s", authorizationTypeBearer, token)
				request.Header.Add(authorizationHeaderKey, authorization)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			url := "/auth"
			
			server := newTestServer(t, nil)
			server.router.GET(url, authMiddleware(server.tokenMaker), func(ctx *gin.Context){
				ctx.JSON(http.StatusOK, gin.H{})
			});
			
			request, err := http.NewRequest(http.MethodGet, url, nil) 
			require.NoError(t, err)

			tc.setAuth(request, server.tokenMaker)
			
			recorder := httptest.NewRecorder()
			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)				
		})
	}

}


