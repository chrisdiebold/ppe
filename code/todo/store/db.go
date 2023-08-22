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
		return nil, fmt.Errorf("could not establish connection to %s", url)
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

func DeleteTodo(db *sql.DB, id int) error {
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

func ConstructTodoList(db *sql.DB) {
	query := `
		SELECT id, name, description, completed, createdOn, CompletedOn
		FROM todos;
	`
	if rows, err := db.Query(query); err != nil {
		l := todo.TodoList{}
		var t todo.Todo
		for rows.Scan(&t) {
			l.Add(t)
		}

		for _, todo := range l {
			fmt.Println(todo.Name)
		}
	}

}

// func GetTodo(db *sql.DB, id int) (*todo.Todo, error) {
// 	query := `
// 		SELECT id, name, description, completed, createdOn, CompletedOn
// 		FROM todos WHERE id = ?;
// 	`
// 	rows, err := db.Query(query, id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(rows) == 0 {
// 		// err
// 	}

// 	// construct todolist

// 	// getTodoErr := rows.Scan(rows)

// 	// todoItem, mapErr := mapQueryToTodo(rows)

// 	// if mapErr != nil {
// 	// 	return nil, getTodoErr
// 	// }

// 	// return &todoItem, nil
// }

// func mapQueryToTodo(rows *sql.Rows) (*todo.Todo, error) todo.TodoList {
// 	newTodo := todo.Todo{}
// 	todoList := todo.TodoList{}
// 	for rows.Scan(&newTodo) {
// 		// add to todolist
// 	}
// 	// for rows.Scan(&newTodo) {

// 	// }
// }
