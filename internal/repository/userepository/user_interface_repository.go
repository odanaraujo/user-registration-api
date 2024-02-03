package userepository

import (
	"database/sql"
	"github.com/odanaraujo/user-api/internal/database/sqlc"
)

func NewUserRepository(db *sql.DB, sqlc *sqlc.Queries) UserRepository {
	return &repository{db: db, q: sqlc}
}

type repository struct {
	db *sql.DB
	q  *sqlc.Queries
}

type UserRepository interface {
	CreateUser() error
}
