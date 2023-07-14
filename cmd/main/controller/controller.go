package controller

import (
	"database/sql"
	"todo/db"

	"github.com/gin-gonic/gin"
)

// GetTodos 	godoc
// @Summary		Get todos
// @Description	Fetch todos from DB
// @Produce		application/json
// @Tags 		todo
// @Success 	200 {object} models.Todo{}
// @Router		/todos [get]
func GetTodos(ctx *gin.Context, database sql.DB) {
	ctx.JSON(200, db.GetTodos(&database))
}
