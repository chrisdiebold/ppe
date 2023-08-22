package store_test

import (
	"fmt"
	"testing"

	"github.com/chrisdiebold/todo"
	"github.com/chrisdiebold/todo/store"
)

var dbUrl = "todos.db"

var testName = "first todo"
var description = "first Todo description"

// SCRATCH:

// func FakeTest(t *testing.T) {
// 	testCases := []struct {
// 		name
// 		string
// 		file
// 		string
// 		ext
// 		string
// 		minSize  int64
// 		expected bool
// 	}{
// 		{"FilterNoExtension", "testdata/dir.log", "", 0, false},
// 		{"FilterExtensionMatch", "testdata/dir.log", ".log", 0, false},
// 		{"FilterExtensionNoMatch", "testdata/dir.log", ".sh", 0, true},
// 		{"FilterExtensionSizeMatch", "testdata/dir.log", ".log", 10, false},
// 		{"FilterExtensionSizeNoMatch", "testdata/dir.log", ".log", 20, true},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			info, err := os.Stat(tc.file)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			f := filterOut(tc.file, tc.ext, tc.minSize, info)
// 			if f != tc.expected {
// 				t.Errorf("Expected '%t', got '%t' instead\n", tc.expected, f)
// 			}
// 		})
// 	}
// }

// func TestStore(t *testing.T) {
// 	type testCase struct {
// 		input todo.TodoList
// 		want  string
// 	}

// 	cases := []testCase{
// 		{
// 			input: kksks,
// 			want: slsls,
// 		},
// 	}

// 	for i, tc := range cases {
// 		got := something(tc.input)
// 		if tc.want != got {
// 			t.Errorf("Fail: Test Case %d failed", i+1)
// 		}
// 	}
// }

func TestGetConnection(t *testing.T) {
	// t.Parallel()
	db, err := store.GetDbConnection(dbUrl)
	defer store.CloseConnection(db)
	if err != nil {
		t.Error("Expected a connection to connect")
	}
	// db.Close()
	// if db.Stats().InUse != 1 {
	// 	fmt.Println(db.Stats().InUse)
	// 	t.Errorf("Failed to open only a single connection")
	// }
}

// if err := dec.Decode(&val); err != nil {
//
//	if serr, ok := err.(*json.SyntaxError); ok {
//	    line, col := findLine(f, serr.Offset)
//	    return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
//	}
//
// return err
// }
func TestSetupTables(t *testing.T) {
	// t.Parallel()
	db, err := store.GetDbConnection(dbUrl)
	// defer store.CloseConnection(db)

	if err != nil {
		t.Error("Expected a connection to connect")
	}
	setupErr := store.SetUpTables(db)

	if setupErr != nil {
		t.Error(setupErr.Error())
	}
}

func TestCloseConnection(t *testing.T) {
	// t.Parallel()
	db, err := store.GetDbConnection(dbUrl)
	// defer store.CloseConnection(db)

	if err != nil {
		t.Error(err.Error())
	}

	closeErr := store.CloseConnection(db)

	if closeErr != nil {
		t.Error(closeErr.Error())
	}
}

func TestClearTodoTable(t *testing.T) {
	// t.Parallel()
	db, _ := store.GetDbConnection(dbUrl)
	// defer store.CloseConnection(db)

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
	// defer store.CloseConnection(db)
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

func TestDeleteTodo(t *testing.T) {
	// t.Parallel()
	todoList := todo.TodoList{}
	todoList.Add(testName, description)

	db, _ := store.GetDbConnection(dbUrl)
	defer store.CloseConnection(db)
	store.SetUpTables(db)

	todoItem, _ := todoList.Get(1)

	store.AddTodo(db, todoItem)

	err := store.DeleteTodo(db, 1)
	if err != nil {
		t.Error("Could not delete a table that did exist.")
	}
}

func TestConstructTodoList(t *testing.T) {
	// store.SeedDatabase()
}

func TestGetTodo(t *testing.T) {
	// t.Parallel()
	todoList := todo.TodoList{}
	todoList.Add(testName, description)

	db, _ := store.GetDbConnection(dbUrl)
	// defer store.CloseConnection(db)
	store.SetUpTables(db)

	todoItem, _ := todoList.Get(1)

	store.AddTodo(db, todoItem)

	todoItem, err := store.GetTodo(db, 1)

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(todoItem.Name)
}
