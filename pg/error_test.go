package pg

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "test"
)

// InitializeDB initializes the database by creating necessary tables
func InitializeDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Create necessary tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(10),
			age smallint,
			created_at TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS some_table (
		    id SERIAL PRIMARY KEY,
		    user_id INT REFERENCES users(id),
		    unique_field VARCHAR(30) UNIQUE,
		    not_null_field VARCHAR(30) NOT NULL
		);
		INSERT INTO some_table (unique_field, not_null_field) VALUES ('Duplicate Name', 'not null value');
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	return db, nil
}

// CleanUpDB drops all tables in the database
func CleanUpDB(db *sql.DB) error {
	_, err := db.Exec(`
		DROP TABLE IF EXISTS some_table;
		DROP TABLE IF EXISTS users;
	`)
	if err != nil {
		return fmt.Errorf("failed to drop tables: %v", err)
	}

	return nil
}

func TestDatabaseErrors(t *testing.T) {
	// Establish a connection to the PostgreSQL database
	db, err := InitializeDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		err := CleanUpDB(db)
		if err != nil {
			panic(err)
		}
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}()

	// Attempt to execute queries that generate various errors
	tests := []struct {
		name     string
		query    string
		expected string // Expected SQL state error code
	}{
		{"String Data Right Truncation", "INSERT INTO users (name) VALUES ('some_really_long_string_that_exceeds_column_length')", "22001"},
		{"Numeric Value Out of Range", "INSERT INTO users (age) VALUES (999999)", "22003"},
		{"Invalid Datetime Format", "INSERT INTO users (created_at) VALUES ('invalid_datetime_value')", "22007"},
		{"Foreign Key Violation", "INSERT INTO some_table (user_id, not_null_field) VALUES (9999, 'Child Name')", "23503"},
		{"Unique Violation", "INSERT INTO some_table (unique_field, not_null_field) VALUES ('Duplicate Name', 'not null value')", "23505"},
		{"Not Null Violation", "INSERT INTO some_table (unique_field) VALUES ('Name')", "23502"},
		{"Syntax Error", "INSERT INTO some_table VALUES", "42601"},
		//{"Insufficient Privilege", "DROP TABLE some_table", "42501"},
		{"Undefined Column", "SELECT non_existing_column FROM some_table", "42703"},
		//{"Connection Failure", "SELECT * FROM some_table", "08006"},
		//{"SQL Client Unable to Establish SQL Connection", "SELECT * FROM non_existing_table", "08001"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := db.Exec(test.query)
			if err == nil {
				t.Fatalf("Expected error, but got none")
			}

			pqErr, ok := err.(*pq.Error)
			if !ok {
				t.Fatalf("Expected a pq.Error, but got: %v", err)
			}

			if string(pqErr.Code) != test.expected {
				t.Fatalf("Expected SQL state error code %s, but got %s", test.expected, pqErr.Code)
			}
		})
	}
}
