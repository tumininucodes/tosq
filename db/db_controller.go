package db

import (
	"database/sql"
	f "fmt"
	"todo/db/models"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() *sql.DB {

	db, error := sql.Open("mysql", "root:alade2001@tcp(localhost:3306)/testdb")
	if error != nil {
		f.Println("error validating sql.Open arguments")
		panic(error.Error())
	} else {
		f.Println("Successfully opened")
	}

	error = db.Ping()

	if error != nil {
		f.Println("error verifying connection with db.Ping") 
		panic(error.Error())
	} else {
		f.Println("db alive")
	}

	return db
}


func GetTodos(db *sql.DB) []models.Todo {

	todos := []models.Todo{}


	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		f.Println("Error executing query:", err)
		return []models.Todo{}
	}
	defer rows.Close()


	for rows.Next() {
		var todo models.Todo


		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body)
		if err != nil {
			f.Println("Error retrieving data:", err)
			return []models.Todo{}
		}

		todos = append(todos, todo)

		f.Println("id:", todo.Id, "body:", todo.Body, "createdAt:", todo.CreatedAt)
	}
	if err := rows.Err(); err != nil {
		f.Println("Error retrieving data:", err)
		return []models.Todo{}
	}

	return todos
}


// insert, err := db.Query("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('2', 'Ben', 'Ford');")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// f.Println("Successful connection to database")

	
