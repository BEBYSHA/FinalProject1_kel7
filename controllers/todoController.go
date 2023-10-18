package controllers

import (
	"net/http"

	"github.com/Bobby-P-dev/FinalProject1_kel7/initializers"
	"github.com/Bobby-P-dev/FinalProject1_kel7/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo struct {
		Name        string
		Description string
	}

	c.Bind(&todo)
	todos := models.Todo{Name: todo.Name, Description: todo.Description}
	result := initializers.DB.Create(&todos)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"message": "succes create data",
		"todos":   todos,
	})

}

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo

	// Querying to find todo datas.
	err := initializers.DB.Find(&todos)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating http response
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success Get data",
		"data":    todos,
	})

}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo struct {
		Name        string
		Description string
	}

	c.Bind(&todo)

	var update models.Todo
	initializers.DB.First(&update, id)

	initializers.DB.Model(&update).Updates(models.Todo{
		Name: todo.Name, Description: todo.Description})

	c.JSON(200, gin.H{
		"update": update,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Todo{}, id)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success delete data",
		"data":    id,
	})
}
