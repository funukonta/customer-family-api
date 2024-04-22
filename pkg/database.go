package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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

	err = createTable(db)
	if err != nil {
		log.Panic("migration failed", err)
	}

	return db
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS nationality (
		nationality_id SERIAL PRIMARY KEY,
		nationality_name VARCHAR(50) NOT NULL,
		nationality_code CHAR(2) NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS customer (
		cst_id SERIAL PRIMARY KEY,
		nationality_id INT NOT NULL,
		cst_name CHAR(50) NOT NULL,
		cst_dob DATE NOT NULL,
		cst_phoneNum VARCHAR(20) NOT NULL,
		cst_email VARCHAR(50) NOT NULL,
		FOREIGN KEY (nationality_id) REFERENCES nationality(nationality_id)
	);
	
	CREATE TABLE IF NOT EXISTS family_list (
		fl_id SERIAL PRIMARY KEY,
		cst_id INT NOT NULL,
		fl_relation VARCHAR(50) NOT NULL,
		fl_name VARCHAR(50) NOT NULL,
		fl_dob DATE NOT NULL,
		FOREIGN KEY (cst_id) REFERENCES customer(cst_id)
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	query = `INSERT INTO nationality (nationality_name, nationality_code) VALUES
	('United States', 'US'),
	('United Kingdom', 'GB'),
	('Canada', 'CA'),
	('Australia', 'AU'),
	('France', 'FR'),
	('Indonesia', 'ID');`

	_, err = db.Exec(query) // ignore error
	if err != nil {
		return nil
	}

	return nil
}
