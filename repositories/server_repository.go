package repositories

import "database/sql"

type ServerRepositoryInterface interface{}

type ServerRepositoryImplementation struct {
	db *sql.DB
}

func NewServerRepositoryImplementation(db *sql.DB) *ServerRepositoryImplementation {
	return &ServerRepositoryImplementation{
		db: db,
	}
}
