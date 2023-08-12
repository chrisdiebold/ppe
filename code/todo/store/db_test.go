package store_test

import (
	"testing"

	"github.com/chrisdiebold/todo/store"
)

var dbUrl = "code/todo/todos.db"

func TestGetConnection(t *testing.T) {
	t.Parallel()
	_, err := store.GetDbConnection(dbUrl)

	if err != nil {
		t.Error("Expected a connection to connect")
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
