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

	


func main(){


	config, err := utils.LoadConfig() 
	if err != nil {
		log.Fatal("cannot load config !")
	}

	con, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database !")
	}


	maker, err := token.NewPasetoMaker(config.TokenSymtricKey)
	if err != nil {
		log.Fatal("cannot create token maker :", err)
	}

	store := db.NewStore(con)
	server := api.NewServer(store, maker, config)
	
	server.Start("0.0.0.0:4000")
}
