package main

//go:generate sqlboiler --wipe --struct-tag-casing camel psql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/uhey22e/sqlboiler-tutor/models"
	"github.com/uhey22e/sqlboiler-tutor/todo"
	"github.com/volatiletech/null"
)

func main() {
	// connect to db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "sample_db",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// validate whether or not the connection string was correct
	if err := db.Ping(); err != nil {
		panic(err)
	}

	todoApp := &todo.App{
		db: db,
	}
}
