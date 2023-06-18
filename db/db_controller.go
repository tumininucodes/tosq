package db

import (
	"database/sql"
	"fmt"
	"time"
	"todo/db/models"

	"github.com/gin-gonic/gin"
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

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt)
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

func CreateTodo(db *sql.DB, todo *models.Todo) []models.Todo {

	todos := []models.Todo{}

	timeStr := "2023-06-16 03:12:20.350418 +0000 UTC"
	layoutIn := "2006-01-02 15:04:05.000000 -0700 MST"
	layoutOut := "2006-01-02T15:04:05Z"

	parsedTime, err := time.Parse(layoutIn, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		panic(err.Error())
	}

	formattedTime := parsedTime.Format(layoutOut)

	result, err := db.Exec("INSERT INTO todo (title, description, createdAt) VALUES (?, ?, ?)", todo.Title, todo.Description, formattedTime)
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT id, title, description, createdAt FROM todo WHERE id = ?", lastInsertID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}

	return todos

}

func UpdateTodo(db *sql.DB, todo *models.Todo) *models.Todo {
	timeStr := "2023-06-16 03:12:20.350418 +0000 UTC"
	layoutIn := "2006-01-02 15:04:05.000000 -0700 MST"
	layoutOut := "2006-01-02T15:04:05Z"

	parsedTime, err := time.Parse(layoutIn, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		panic(err.Error())
	}

	formattedTime := parsedTime.Format(layoutOut)
	todo.CreatedAt = formattedTime

	_, err = db.Exec("UPDATE todo SET title = ?, description = ?, createdAt = ? WHERE id = ?", todo.Title, todo.Description, todo.CreatedAt, todo.Id)
	if err != nil {
		panic(err.Error())
	}

	return todo
}

func DeleteTodo(db *sql.DB, id string) *gin.H {
	_, err := db.Exec("DELETE FROM todo where id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	return &gin.H{"status":"Delete operation successful"}
}
