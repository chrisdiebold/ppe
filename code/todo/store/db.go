package store

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func GetDbConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "todo/todos.db")

	if err != nil {
		return nil, fmt.Errorf("Could not establish connection to %s", url)
	}
	return db, nil
}

// TODO: This method does not seem to be very useful
func CloseConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}