package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct { //Требования
	ID     string `json:"id"`   //GET запрос на localhost:8080/tasks возвращает hello task
	Name   string `json:"name"` //POST запрос на localhost:8080/tasks передает json c полем tasks и записывать его содержиме в переменную
	Status string `json:"status"`
}

type TaskRequest struct {
	Name string `json:"name"`
}

var tasks = []Task{}
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
	globalTask = req.Name

	task := Task{
		ID:     uuid.NewString(),
		Name:   req.Name,
		Status: "active",
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
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
