package db

import (
	"database/sql"
	"fmt"
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
		return []models.Todo{}
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body)
		if err != nil {
			fmt.Println("Error retrieving data:", err)
			return []models.Todo{}
		}

		todos = append(todos, todo)

		fmt.Println("id:", todo.Id, "body:", todo.Body, "createdAt:", todo.CreatedAt)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
		return []models.Todo{}
	}

	return todos
}

func CreateTodo(db *sql.DB) []models.Todo {

	todos := []models.Todo{}


	result, err := db.Exec("INSERT INTO todo (title, description) VALUES ('edkndkn', 'wojqew')")
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT id, title, description FROM todo WHERE id=?", lastInsertID)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}

	return todos

}
