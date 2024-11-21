package repositories

import "database/sql"

type ExampleRepositoryInterface interface{}

type ExampleRepositoryImplementation struct {
	db *sql.DB
}

func NewExampleRepositoryImplementation(db *sql.DB) *ExampleRepositoryImplementation {
	return &ExampleRepositoryImplementation{
		db: db,
	}
}
