package main

import (
	"database/sql"
	"log"

	"github.com/brkss/go-auth/api"
	db "github.com/brkss/go-auth/db/sqlc"
	"github.com/brkss/go-auth/token"
	"github.com/brkss/go-auth/utils"
	_ "github.com/lib/pq"
)

	
const (
	DBDriver = "postgres"
	DBSource = "postgres://root:root@localhost:5432/auth?sslmode=disable"
)

func main(){

	con, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to database !")
	}

	maker, err := token.NewPasetoMaker(utils.RandomString(32))
	if err != nil {
		log.Fatal("cannot create token maker :", err)
	}

	store := db.NewStore(con)
	server := api.NewServer(store, maker)
	
	server.Start("0.0.0.0:4000")
}
