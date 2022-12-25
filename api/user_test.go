package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/brkss/go-auth/db/mock"
	db "github.com/brkss/go-auth/db/sqlc"
	"github.com/brkss/go-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParams struct {
	arg 		db.CreateUserParams
	password 	string
}

func (e eqCreateUserParams)Matches(x interface{}) bool {

	// cast interface 
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}
	// VERIFY PASSWORD 
	err := utils.VerifyPassword(arg.Password, e.password)
	if err != nil {
		return false
	}
	e.arg.Password = arg.Password
	e.arg.ID = arg.ID
	return true
}

func (e eqCreateUserParams)String() string {
	return fmt.Sprintf("password %v matches %v", e.password, e.arg.Password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParams{
		arg: arg,
		password: password,
	}
}

func CreateUser(t *testing.T)(db.User, string){
	password := utils.RandomString(10)
	hash, err := utils.HashPassword(password)
	require.NoError(t, err)
	user := db.User{
		ID: uuid.New().String(),
		Username: utils.RandomName(),
		Email: utils.RandomEmail(),
		Name: utils.RandomName(),
		Password: hash,
		CreatedAt: time.Now(),
	}
	return user, password
}

func TestRegisterUser(t *testing.T){

	user, password := CreateUser(t)

	testCases := []struct{
		name			string
		body 			gin.H
		buildStabs 		func(store *mockdb.MockStore)
		checkResponse 	func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name": user.Name,
				"username": user.Username,
				"email": user.Email,
				"password": password,
			},
			buildStabs: func(store *mockdb.MockStore){
				arg := db.CreateUserParams{
					ID: user.ID,
					Username: user.Username,
					Email: user.Email,
					Password: user.Password,
					Name: user.Name,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"name": user.Name,
				"email": user.Email,
				"password": password,
			},
			buildStabs: func(store *mockdb.MockStore){
				
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name": user.Name,
				"username": user.Username,
				"email": user.Email,
				"password": password,
			},
			buildStabs: func(store *mockdb.MockStore){
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError , recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStabs(store)

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/register"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))  
			require.NoError(t, err)

			server := newTestServer(t, store) 
			recorder := httptest.NewRecorder()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
