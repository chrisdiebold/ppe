package store_test

import (
	"testing"

	"github.com/chrisdiebold/todo"
	"github.com/chrisdiebold/todo/store"
)

var dbUrl = "todos.db"

func TestGetConnection(t *testing.T) {
	t.Parallel()
	db, err := store.GetDbConnection(dbUrl)

	if err != nil {
		t.Error("Expected a connection to connect")
	}
	db.Close()
}

func TestSetupTables(t *testing.T) {
	t.Parallel()
	db, err := store.GetDbConnection(dbUrl)

	if err != nil {
		t.Error("Expected a connection to connect")
	}
	setupErr := store.SetUpTables(db)

	if setupErr != nil {
		t.Error(setupErr.Error())
	}
}

func TestCloseConnection(t *testing.T) {
	t.Parallel()
	db, err := store.GetDbConnection(dbUrl)

	if err != nil {
		t.Error(err.Error())
	}

	closeErr := store.CloseConnection(db)

	if closeErr != nil {
		t.Error(closeErr.Error())
	}
}

func TestClearTodoTable(t *testing.T) {
	t.Parallel()
	db, _ := store.GetDbConnection(dbUrl)

	store.SetUpTables(db)

	err := store.ClearTodoTable(db)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddTodo(t *testing.T) {
	todoList := todo.TodoList{}
	todoList.Add("First task", "Do tickets for Chris!")

	db, _ := store.GetDbConnection(dbUrl)
	defer store.CloseConnection(db)
	store.SetUpTables(db)

	todoItem, _ := todoList.Get(1)
	affectedRows, err := store.AddTodo(db, todoItem)

	if err != nil {
		t.Error(err.Error())
	}

	if affectedRows != 1 {
		t.Error("Did not affect expected number of rows")
	}

}
