package store

import (
	"database/sql"
	"fmt"

	"github.com/chrisdiebold/todo"
)

func SeedDatabase(database *sql.DB) error {
	var errToReturn error

	err := ClearTodoTable(database)
	if err != nil {
		errToReturn = err
	}
	err2 := SetUpTables(database)
	if err2 != nil {
		errToReturn = err
	}

	for i := 0; i < 5; i++ {
		t := fmt.Sprintf("A test todo %d", i)
		todo := todo.NewTodo(t, t)
		AddTodo(database, todo)
	}

	return errToReturn
}
