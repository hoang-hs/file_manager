package repositories

import "database/sql"

type baseRepository struct {
	db *sql.DB
}

func NewBaseRepository(db *sql.DB) *baseRepository {
	return &baseRepository{
		db: db,
	}
}
