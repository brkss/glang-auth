package api

import (
	"os"
	"testing"

	db "github.com/brkss/go-auth/db/sqlc"
	"github.com/brkss/go-auth/token"
	"github.com/brkss/go-auth/utils"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store)(*Server){
	
	maker, err := token.NewPasetoMaker(utils.RandomString(32)) 
	require.NoError(t, err)
	
	server := NewServer(store, maker)
	return (server)
}

func TestMain(m *testing.M) {



	os.Exit(m.Run())

}
