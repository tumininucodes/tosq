package main

import (
	"fmt"
	"todo/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func main() {


	fmt.Println("start of application")
	server := gin.Default()

	database := db.OpenDB
	defer database().Close()

	server.GET("/todos", func(ctx *gin.Context) {
		ctx.JSON(200, db.GetTodos(database()))
	})

	server.Run(":8080")
	

}