package token

import (
	"testing"
	"time"

	"github.com/brkss/go-auth/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidToken(t *testing.T) {

	id := uuid.New().String()

	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(id, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, payload.UserId, id)
	require.WithinDuration(t, payload.IssuedAt, time.Now(), time.Second)
	require.WithinDuration(t, payload.ExpiredAt, time.Now().Add(time.Minute), time.Second)
}

func TestExpiredToken(t *testing.T) {
	id := uuid.New().String()

	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(id, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	_, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
}

func TestInvalidToken(t *testing.T) {

	id := uuid.New().String()

	maker1, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	maker2, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, err := maker1.CreateToken(id, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	_, err = maker2.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
}

func TestInvalidSymtrictKey(t *testing.T) {

	_, err := NewPasetoMaker("aaa")
	require.Error(t, err)
}
