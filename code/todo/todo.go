package todo

import (
	"fmt"
	"time"
)

// Todolist is a slice that represents a todo list
type TodoList []todo

type todo struct {
	Id          int
	Name        string
	Description string
	Completed   bool
	CreatedOn   time.Time
	CompletedOn time.Time
}

func (l *TodoList) Add(task string, desc string) {

	todoItem := todo{
		Id:          0,
		Name:        task,
		Description: desc,
		Completed:   false,
		CreatedOn:   time.Now(),
		CompletedOn: time.Time{},
	}
	*l = append(*l, todoItem)

}

// Adjust for 1 based lists: *l = append(ls[:i-1], ls[i:]...)
// TODO: Complete(int)
func (l *TodoList) Complete(item int) error {
	ls := *l

	if item <= 0 || item > len(ls) {
		return fmt.Errorf("item %d does not exist", item)
	}

	ls[item-1].Completed = true
	ls[item-1].CompletedOn = time.Now()

	return nil
}

// Delete removes the numbered element from the todo list
// If the list is empty or the user get an item out of bounds will give an error
func (l *TodoList) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// TODO: GET(item int) // Get the Todo Item.
func (l * TodoList) Get(i int) (*todo, error) {
	ls := *l

	if i <= 0 || i > len(ls) {
		return nil, fmt.Errorf("item %d does not exist", i)
	}

	return &ls[i-1], nil
}

func NewTodo(name, description string) *todo {
	return &todo{
		Name:        name,
		Description: description,
		CreatedOn:   time.Now(),
		CompletedOn: time.Time{},
		Completed:   false,
	}
}
