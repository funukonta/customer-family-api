package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectPostgres() *sql.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, ssl)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicln("error connStr", err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic("connection db not valid", err)
	}

	return db
}
