package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"todo/db"
	"todo/db/models"

	"github.com/gin-gonic/gin"
)

// GetTodos 	godoc
// @Summary		Get todos
// @Description	Fetch todos from DB
// @Produce		application/json
// @Tags 		todo
// @Success 	200 {object} models.Todo{}
// @Router		/todos [get]
func GetTodos(ctx *gin.Context, database *sql.DB) {
	ctx.JSON(200, db.GetTodos(database))
}

// CreateTodo 	godoc
// @Summary		Create todo
// @Description	Add a todo to the DB
// @Produce		application/json
// @Param todo body models.Todo{} true "Todo object that is to be created"
// @Tags 		todo
// @Success 	200 {object} models.Todo{}
// @Router		/todo [post]
func CreateTodo(ctx *gin.Context, database *sql.DB) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, db.CreateTodo(database, &todo))
}

// DeleteTodo 	godoc
// @Summary		Delete todo
// @Description	Delete a todo from the DB
// @Produce		application/json
// @Param 		id path int true "Todo ID"
// @Tags 		todo
// @Success 	200 {object} gin.H{}
// @Router		/todo/:id [delete]
func DeleteTodo(ctx *gin.Context, database *sql.DB, id string) {
	ctx.JSON(200, db.DeleteTodo(database, id))
}

// UpdateTodo 	godoc
// @Summary		Update todo
// @Description	Update a todo in the DB
// @Produce		application/json
// @Param todo body models.Todo{} true "Todo object that is to be updated"
// @Tags 		todo
// @Success 	200 {object} models.Todo{}
// @Router		/todo/:id [put]
func UpdateTodo(idString string, ctx *gin.Context, database *sql.DB) {
	todo := &models.Todo{}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "wrong id passed"})
	}
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		panic(err.Error())
	}
	todo.Id = int64(id)
	ctx.JSON(200, db.UpdateTodo(database, todo))
}
