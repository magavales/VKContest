package database

import (
	"VK/pkg/database/tables"
	"github.com/jackc/pgx"
	"log"
)

type Database struct {
	Pool  *pgx.Conn
	Users tables.UsersDB
}

func (db *Database) Connect() {
	config := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "1703",
	}

	var err error
	db.Pool, err = pgx.Connect(config)
	if err != nil {
		log.Printf("I can't connect to database: %s\n", err)
	}
}
