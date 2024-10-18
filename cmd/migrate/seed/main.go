package main

import (
	"log"

	"github.com/filipdoescode/go-social/internal/db"
	"github.com/filipdoescode/go-social/internal/store"
)

func main() {
	dbaddr := "postgres://admin:password@localhost/social?sslmode=disable"

	conn, err := db.New(dbaddr, 30, 30, "15m")

	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(conn)

	db.Seed(store)
}
