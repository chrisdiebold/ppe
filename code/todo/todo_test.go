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

func TestComplete(t *testing.T) {
	list := todo.TodoList{}

	title1 := "First Todo"
	desc1 := "First desc"
	title2 := "Second Todo"
	desc2 := "Second desc"

	list.Add(title1, desc1)
	list.Add(title2, desc2)

	list.Complete(2)

	if list[1].Completed != true {
		t.Error("Expected true, got false instead")
	}
}
