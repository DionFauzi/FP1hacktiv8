package routes

import (
	"FP1hacktiv8/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRouters(router *gin.Engine) {
	router.GET("/todos", controllers.GetAllTodos)
	router.GET("/todos/:id", controllers.GetAllTodosByID)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/update/:id", controllers.UpdateTodo)
	router.DELETE("/todos/delete/:id", controllers.DeleteTodo)
}
