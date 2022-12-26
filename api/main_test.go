package api

import (
	"os"
	"testing"
	"time"

	db "github.com/brkss/go-auth/db/sqlc"
	"github.com/brkss/go-auth/token"
	"github.com/brkss/go-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const (
	DBDriver = "postgres"
	DBSource = "postgres://root:root@localhost:5432/auth?sslmode=disable"
)

func newTestServer(t *testing.T, store db.Store)(*Server){
	
	maker, err := token.NewPasetoMaker(utils.RandomString(32)) 
	require.NoError(t, err)

	config := &utils.Config{
		DBSource: DBSource,
		DBDriver: DBDriver,
		TokenSymtricKey: utils.RandomString(32),
		TokenDuration: time.Minute * 15,
	}

	server := NewServer(store, maker, config)
	return (server)
}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())

}
