package db

import "database/sql"

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(con *sql.DB) Store {
	return &SQLStore{
		db: con,
		Queries: New(con),
	}
}
