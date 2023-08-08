package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func (todo *Todo) getName() {
	return todo.name
}

func main() {
	db, err := sql.Open("sqlite", "todos.db")
	if err != nil {
		panic("Could not connect to the database")
	}

	ping := db.Ping()
	if ping == nil {
		fmt.Println("it works!")
		// do not forget to remove all tables from db prior to staring....
		// TODO: Create the db scheme
		// TODO: Insert 4 todos in the db as a transaction
		// TODO: Query those todos and dump them to the consol
		// TODO: In todo model create method for:
		// TODO: completing a todo
		// TODO: deleting a todo
		// TODO: updating a todo name

		// TODO: stretch put all of this in a seed.go file and use it in main.

	}

}

func SetupTables(data *sql.DB) {
	// https://sqlite.org/lang_createtable.html
	// var todoSchema = """
	// 	CREATE TABLE IF NOT EXISTS todo (
	// 		id INTEGER PRIMARY KEY ASC,
	// 		name VARCHAR ( 200 ) UNIQUE NOT NULL,
	// 		completed bool NOT NULL,
	// 		createdAt DATETIME NOT NULL,
	// 		completedAt DATETIME
	// 	)
	// """

}
