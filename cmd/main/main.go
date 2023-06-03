package main

import (
	"fmt"
	"todo/db"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {

	fmt.Println("start of application")

	database := db.OpenDB


	defer database().Close()

	todos := db.GetTodos(database())

	fmt.Println(todos)
	

}