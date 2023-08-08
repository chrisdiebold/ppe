package todo_test

import (
	"testing"

	"github.com/chrisdiebold/todo"
)

func TestAdd(t *testing.T) {
	list := todo.TodoList{}

	title := "An important todo"
	d := "learn to code go better"

	list.Add(title, d)

	if list[0].Name != title {
		t.Errorf("Expected %q, got %q instead.", title, list[0].Name)
	}

	if list[0].Description != d {
		t.Errorf("Expected %q, got %q instead.", d, list[0].Description)
	}
}
