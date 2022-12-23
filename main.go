package main

import (
	"database/sql"
	"log"

	"github.com/brkss/go-auth/api"
	db "github.com/brkss/go-auth/db/sqlc"
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

	store := db.NewStore(con)
	server := api.NewServer(store)
	
	server.Start("0.0.0.0:4000")
}
