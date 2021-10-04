package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DB_URL"))

	DB = db

	return DB, err
}
