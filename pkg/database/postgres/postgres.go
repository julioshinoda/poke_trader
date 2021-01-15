package postgres

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4"
	"gocloud.dev/postgres"
	_ "gocloud.dev/postgres/gcppostgres"
	//"os"
)

func GetConnection() (*sql.DB, error) {
	conn, err := postgres.Open(context.Background(), os.Getenv("DATABASE_URL")) //"gcppostgres://nucleus-dev:!oeowNohx213ycr26nEk87ZZ@nucleus-dev-297217/europe-west1/nucleus-dev/nucleus"
	conn.SetMaxOpenConns(12)
	conn.SetMaxIdleConns(12)
	conn.SetConnMaxLifetime(2 * time.Minute)
	return conn, err
}
