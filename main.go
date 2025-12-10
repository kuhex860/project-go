package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Task string `json:"task"` // Изменили Name на Task для соответствия ТЗ
}

var tasks = []Task{} // Список задач
var globalTask string = "Задача не установлена"

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello, task: "+globalTask)
}

func postTask(c echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	globalTask = req.Task

	newTask := Task{
		ID:     uuid.NewString(),
		Name:   req.Task,
		Status: "active",
	}

	tasks = append(tasks, newTask)

	return c.JSON(http.StatusCreated, newTask)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	for i, Task := range tasks {
		if Task.ID == id {
			tasks[i].Name = req.Task
			tasks[i].Status = "active"
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	for i, Task := range tasks {
		if Task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

//func getHello(c echo.Context) error {
//	return c.String(http.StatusOK, "hello, task: "+globalTask)
//}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	//	e.GET("/", getHello)
	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTask)
	e.PATCH("/tasks/:id", patchTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Start("localhost:8080")
}
