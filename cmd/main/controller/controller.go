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

func CreateTodo(ctx *gin.Context, database *sql.DB) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, db.CreateTodo(database, &todo))
}

func DeleteTodo(ctx *gin.Context, database *sql.DB, id string) {
	ctx.JSON(200, db.DeleteTodo(database, id))
}

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
