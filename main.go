package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var tasks = []Task{}
var nextID = 1

func main() {

	r := gin.Default()

	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTaskByID)
	r.PUT("/tasks/:id", deleteTask)

	r.Run(":8080")
}

func createTask(c *gin.Context){
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err !=
	nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = nextID
	nextID++

	tasks = append(tasks, newTask)

	//O que acontece aqui
	//Recebe JSON
	// Converte pra struct
	// Gera ID
	// Salva no slice
	// Retorna resposta

	c.JSON(http.StatusCreated, newTask)
}

func getTasks(c *gin.Context){
	c.JSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	for _, t := range tasks {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func updateTask(c *gin.Context){

	id, _ := strconv.Atoi(c.Param("id"))
	var update Task

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Title = update.Title
			tasks[i].Description = update.Description
			tasks[i].Done = update.Done

			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func deleteTask(c *gin.Context){

	id, _ := strconv.Atoi(c.Param("id"))

	for i, t := range tasks {
		if t.ID == id {

			tasks = append(tasks[:i], tasks[1+i:]...)

			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}