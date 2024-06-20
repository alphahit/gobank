package main

import (
	"database/sql"

	_ "github.com/lib/pq" // Import the postgres driver anonymously to initialize it, enabling database/sql to use PostgreSQL.
)

// Storage defines the interface for account storage operations. This abstraction allows for easy testing and future storage implementations.
type Storage interface {
	CreateAccount(*Account) error     // CreateAccount inserts a new account into the store.
	DeleteAccount(int) error          // DeleteAccount removes an account by its ID.
	UpdateAccount(*Account) error     // UpdateAccount modifies an existing account.
	GetAccount(int) (*Account, error) // GetAccount retrieves an account by its ID.
}

// PostGreStore holds the connection details to the PostgreSQL database.
type PostGreStore struct {
	db *sql.DB // db represents a pool of zero or more underlying connections to the database.
}

// NewPostgresStore initializes a new PostGreStore with a connection to a PostgreSQL database.
// It returns the newly created store or an error if the connection or setup fails.
func NewPostgresStore() (*PostGreStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable" // Connection string to connect to the PostgreSQL database.
	db, err := sql.Open("postgres", connStr)                                   // Open a new database connection.

	if err != nil {
		return nil, err // Return an error if the database connection cannot be established.
	}

	if err := db.Ping(); err != nil {
		return nil, err // Return an error if the connection to the database cannot be verified.
	}

	return &PostGreStore{
		db: db, // Initialize PostGreStore with the established database connection.
	}, nil
}
