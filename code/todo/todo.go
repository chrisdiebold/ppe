package todo

import (
	"time"
)

type TodoList []todo

type todo struct {
	Id          int
	Name        string
	Description string
	Completed   bool
	CreatedOn   time.Time
	CompletedOn time.Time
}

// primes := [6]int{2, 3, 5, 7, 11, 13} // This is how to define an array
// primes := []int{2, 3, 5, 7, 11, 13} // slice
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
func (l *TodoList) Complete(itemOrder int) {

	if itemOrder > 0 {
		itemOrder--
	}
	for i := 0; i < len(*l); i++ {

		if i == itemOrder {
			todoItem := *l[i]
			todoItem.Completed = true
			todoItem.CompletedOn = time.Now()
		}
	}
}

// TODO: Delete
// TODO: GET - print the whole list.

func NewTodo(name, description string) *todo {
	return &todo{
		Name:        name,
		Description: description,
		CreatedOn:   time.Now(),
		CompletedOn: time.Time{},
		Completed:   false,
	}
}
