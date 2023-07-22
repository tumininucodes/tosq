package main

import (
	"fmt"
	"todo/cmd/main/controller"
	_ "todo/cmd/main/docs"
	"todo/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Todo API
// @version 1.0
// @description Todo API Documentation. Written in Go. Gin used
// @host localhost:8080
// @BasePath /
func main() {

	fmt.Println("start of application")
	server := gin.Default()

	database := db.OpenDB
	defer database().Close()

	// Add swagger
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get Todos
	server.GET("/todos", func(ctx *gin.Context) {
		controller.GetTodos(ctx, database())
	})

	// Create Todo
	server.POST("/todo", func(ctx *gin.Context) {
		controller.CreateTodo(ctx, database())
	})

	// Delete Todo
	server.DELETE("/todo/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		controller.DeleteTodo(ctx, database(), id)
	})

	// Update Todo
	server.PUT("todo/:id", func(ctx *gin.Context) {
		idString := ctx.Param("id")
		controller.UpdateTodo(idString, ctx, database())
	})

	server.Run(":8080")

}
