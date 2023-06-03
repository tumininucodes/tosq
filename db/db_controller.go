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

	rows, err := db.Query("SELECT * FROM students")
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

	result, err := db.Exec("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('10', 'Ben', 'Ford');")
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(rowsAffected)

	rows, err := db.Query("SELECT `id`, `firstname`, `lastname` FROM `testdb`.`students` WHERE `id`=?", rowsAffected)
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
