package db

import (
	"database/sql"
	"fmt"
	"time"
	"todo/db/models"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() *sql.DB {

	db, error := sql.Open("mysql", "root:alade2001@tcp(localhost:3306)/testdb")
	if error != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(error.Error())
	} else {
		fmt.Println("Successfully opened")
	}

	error = db.Ping()

	if error != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(error.Error())
	} else {
		fmt.Println("db alive")
	}

	return db
}

func GetTodos(db *sql.DB) []models.Todo {

	todos := []models.Todo{}

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		fmt.Println("Error executing query:", err)
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.CreatedAt)
		if err != nil {
			fmt.Println("Error retrieving data:", err)
			panic(err.Error())
		}

		todos = append(todos, todo)
	}
	
	if err := rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
		panic(err.Error())
	}

	return todos
}

func CreateTodo(db *sql.DB) []models.Todo {

	todos := []models.Todo{}

	fmt.Println(time.Now().UTC().String())

	timeStr := "2023-06-16 03:12:20.350418 +0000 UTC"
	layoutIn := "2006-01-02 15:04:05.000000 -0700 MST"
	layoutOut := "2006-01-02T15:04:05Z"

	parsedTime, err := time.Parse(layoutIn, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		panic(err.Error())
	}

	formattedTime := parsedTime.Format(layoutOut)

	result, err := db.Exec("INSERT INTO todo (title, description, createdAt) VALUES (?, ?, ?)", "Hii", "Joe", formattedTime)
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT id, title, description, createdAt FROM todo WHERE id=?", lastInsertID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}

	return todos

}
