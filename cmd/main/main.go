package main

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/db"
	"todo/db/models"
	"github.com/gin-gonic/gin"
	_ "todo/cmd/main/docs"
	_ "github.com/go-sql-driver/mysql"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title Begin
// @version 1.0
// @description Na here we dey
// @host localhost:8080
// @BasePath /
func main() {


	fmt.Println("start of application")
	server := gin.Default()

	database := db.OpenDB
	defer database().Close()

	// Add swagger
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.GET("/todos", func(ctx *gin.Context) {
		ctx.JSON(200, db.GetTodos(database()))
	})

	server.POST("/todo", func(ctx *gin.Context) {
		var todo models.Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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