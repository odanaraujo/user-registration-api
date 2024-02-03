package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/odanaraujo/user-api/config/env"
	"log/slog"
)

func NewDBConnection() (*sql.DB, error) {
	postgresURL := env.Env.DatabaseURL
	db, err := sql.Open("postgres", postgresURL)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	slog.Info("database connected", slog.String("package", "database"))
	return db, nil
}
