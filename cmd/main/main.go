package main

import (
	"fmt"
	"net/http"
	"todo/db"
	"todo/db/models"

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

	server.POST("/todo", func(ctx *gin.Context) {
		var todo models.Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(todo)
		ctx.JSON(201, db.CreateTodo(database(), &todo))
	})


	server.Run(":8080")
	

}