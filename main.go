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

func getHello(c echo.Context) error {
	return c.String(http.StatusOK, "hello, task: "+globalTask)
}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
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

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/", getHello)
	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTask)

	e.Start("localhost:8080")
}
