package controllers

import (
	"FP1hacktiv8/models"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetAllTodos())
}

func GetAllTodosByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	todo, err := models.GetTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data can not created",
		})
		return
	}

	models.AddNewTodo(&newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
	var updateTodo models.Todo

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := c.BindJSON(&updateTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data can not created",
		})
		return
	}
	todo, errUpdate := models.UpdateTodoById(id, &updateTodo)
	if errUpdate != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"msg": "Data" + todo.ToList + "Has Been Update",
	})
}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	todo, err := models.DeleteTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"msg": "Data" + todo.ToList + "Has Been Deleted",
	})
}
