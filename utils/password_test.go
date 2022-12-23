package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestValidPassword(t *testing.T){
	password := RandomString(10)
	
	hash, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = VerifyPassword(hash, password)
	require.NoError(t, err)

}

func TestInvalidPassword(t *testing.T){
	
	password1 := RandomString(10)
	password2 := RandomString(10)

	hash, err := HashPassword(password1)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = VerifyPassword(hash, password2)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
