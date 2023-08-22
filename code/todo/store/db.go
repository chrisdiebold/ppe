package store

import (
	"database/sql"
	"fmt"

	"github.com/chrisdiebold/todo"
	_ "modernc.org/sqlite"
)

func GetDbConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", url)

	if err != nil {
		return nil, fmt.Errorf("Could not establish connection to %s", url)
	}

	return db, nil
}

func SetUpTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS todos(
			id INTEGER PRIMARY KEY,
			name TEXT,
			description TEXT,
			completed BOOLEAN,
			createdOn Date,
			completedOn Date
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func CheckConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// TODO: This method does not seem to be very useful
func CloseConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func ClearTodoTable(db *sql.DB) error {
	query := "DELETE FROM todos"

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func AddTodo(db *sql.DB, todoItem *todo.Todo) (int64, error) {
	query := `
		INSERT INTO todos(
			name,
			description,
			completed,
			createdOn,
			completedOn
		)
		values(?, ?, ?, ?, ?);
	`
	result, err := db.Exec(query, todoItem.Name, todoItem.Description,
		todoItem.Completed, todoItem.CreatedOn, todoItem.CompletedOn)
	if err != nil {
		return 0, err
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, err
}

func DeleteTodo(db *sql.DB, id int) (error) {
	query := `
		DELETE FROM todos WHERE id = ?;
	`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, deleteError := result.RowsAffected()
	if rowsAffected != 1 || deleteError != nil {
		return deleteError
	}
	return nil
}

func GetTodo(db *sql.DB, id int) (*todo.Todo, error) {
	query := `
		SELECT id, name, description, completed, createdOn, CompletedOn
		FROM todos WHERE id = ?;
	`
	rows, err := db.Query(query, id)
	
	if err != nil {
		return nil, err
	}
	todo := todo.Todo{}

	getTodoErr :=  rows.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.CompletedOn)
	if getTodoErr != nil {
		return nil, getTodoErr
	}

	return &todo, nil
}
