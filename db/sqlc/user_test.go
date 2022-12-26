package db

import (
	"context"
	"testing"
	"time"

	"github.com/brkss/go-auth/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)


func CreateUser(t *testing.T) (User){

	arg := CreateUserParams{
		ID: uuid.New().String(),
		Username: utils.RandomName(),
		Email: utils.RandomEmail(),
		Password: utils.RandomString(10),
		Name: utils.RandomName(),
	}
	
	user, err := testQueries.CreateUser(context.Background(), arg) 
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.ID, arg.ID)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.Password, arg.Password)
	require.WithinDuration(t, user.CreatedAt, time.Now(), time.Second)	
	return (user)
}

func TestCreateUser(t *testing.T){
	CreateUser(t)
}


func TestGetUser(t *testing.T){
	user := CreateUser(t)

	gotUser, err := testQueries.GetUser(context.Background(), user.Username);
	require.NoError(t, err)
	require.NotEmpty(t, gotUser)

	require.Equal(t, user.ID, gotUser.ID)
	require.Equal(t, user.Email, gotUser.Email)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.Password, gotUser.Password)
	require.WithinDuration(t, user.CreatedAt, gotUser.CreatedAt, time.Second)
}

func TestMe(t *testing.T){
	user := CreateUser(t)
	
	me, err := testQueries.Me(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, me.Username)
	require.Equal(t, user.Name, me.Name)
	require.Equal(t, user.Email, me.Email)
	require.WithinDuration(t, user.CreatedAt, me.CreatedAt, time.Second)
}
