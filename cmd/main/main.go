package main

import (
	"fmt"
	"net/http"
	"strconv"
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

	server.DELETE("/todo/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		
		ctx.JSON(200, db.DeleteTodo(database(), id))
	})

	server.PUT("todo/:id",  func(ctx *gin.Context) {
		idString := ctx.Param("id")
		todo := &models.Todo{}
		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "wrong id passed"})
		}
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			panic(err.Error())
		}
		todo.Id = int64(id)
		ctx.JSON(200, db.UpdateTodo(database(), todo))
	})


	server.Run(":8080")
	

}