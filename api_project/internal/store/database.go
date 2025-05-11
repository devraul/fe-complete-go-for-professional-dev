package store

import (
	"database/sql"
	"fmt"

	// We must to import it because it'll be used under
	// the hood.
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=postgres passoword=postgres dbname=postgres port=5432 sslmode=disabled")

	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to database...")
	// db.SetMaxOpenConns()
	// db.SetMaxIdleConns()
	// db.SetConnMaxIdleTime()

	return db, err
}
