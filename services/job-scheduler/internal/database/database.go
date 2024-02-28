package database

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Database represents the interface for interacting with a PostgreSQL database.
type Database interface {
	Connect() (*sql.DB, error)
	Close(*sql.DB) error
	Query(*sql.DB, string) (*sql.Rows, error)
	Exec(*sql.DB, string) (sql.Result, error)
}

// PostgreSQL implements the Database interface.
type PostgreSQL struct {
	ConnectionString string
}

// Connect establishes a connection to the PostgreSQL database.
func (p *PostgreSQL) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", p.ConnectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Close closes the connection to the PostgreSQL database.
func (p *PostgreSQL) Close(db *sql.DB) error {
	return db.Close()
}

// Query executes a query on the PostgreSQL database.
func (p *PostgreSQL) Query(db *sql.DB, query string) (*sql.Rows, error) {
	return db.Query(query)
}

// Exec executes a SQL statement on the PostgreSQL database.
func (p *PostgreSQL) Exec(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
