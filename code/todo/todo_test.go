package todo_test

import (
	"testing"

	"github.com/chrisdiebold/todo"
)

func TestAdd(t *testing.T) {
	t.Parallel()
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

	if len(list) != 1 {
		t.Errorf("Expected %d, got %d instead.", 1, len(list))
	}
}

func TestCompleteOnEmptyList(t *testing.T) {
	t.Parallel()
	list := todo.TodoList{}

	err := list.Complete(1)
	if err == nil {
		t.Error("Expected an error to be thrown")
	}
}

func TestComplete(t *testing.T) {
	t.Parallel()

	list := todo.TodoList{}
	list.Add("Learn Go", "understand interfaces")
	list.Add("Learn Go", "understand interfaces")

	if list[0].Completed != false {
		t.Error("Expected true, got false instead")
	}

	list.Complete(1)

	if list[0].Completed != true {
		t.Error("Expected true, got false instead")
	}
}

func TestGet(t* testing.T) {
	t.Parallel()

	list := todo.TodoList{}
	list.Add("Learn Go", "understand interfaces")
	list.Add("Learn Go 2", "understand interfaces")

	_, err := list.Get(10)
	if err == nil {
		t.Error("Expected an error, did not get one instead")
	}

	firstItem, _ := list.Get(1)

	if firstItem.Name != "Learn Go" {
		t.Errorf("Expected %s, got %s instead", list[0].Name, firstItem.Name)
	} 
}

func TestDelete(t *testing.T) {
	l := todo.TodoList{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		l.Add(v, "some desc")
	}

	if len(l) != 3 {
		t.Errorf("Expected list length %d, got %d instead.", 3, len(l))
	}
	if l[0].Name != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Name)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}
	if l[1].Name != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Name)
	}
}
