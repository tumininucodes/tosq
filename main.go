package main

import (
	"database/sql"
	f "fmt"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	f.Print("hello world!")

	db, error := sql.Open("mysql", "root:alade2001@tcp(localhost:3306)/testdb")
	if error != nil {
		f.Println("error validating sql.Open arguments")
		panic(error.Error())
	} else {
		f.Println("Successfully opened")
	}

	defer db.Close()

	error = db.Ping()

	if error != nil {
		f.Println("error verifying connection with db.Ping") 
		panic(error.Error())
	} else {
		f.Println("db alive")
	}

	insert, err := db.Query("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('2', 'Ben', 'Ford');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	f.Println("Successful connection to database")

}