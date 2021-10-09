package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))

	DB = db

	return DB, err
}
