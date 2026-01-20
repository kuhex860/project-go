package main

import (
	"app/internal/TaskService"
	"app/internal/db"
	"app/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}
	taskRepo := TaskService.NewTaskRepository(database)
	taskService := TaskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/tasks", taskHandlers.GetTasks)
	e.POST("/tasks", taskHandlers.PostTask)
	e.PATCH("/tasks/:id", taskHandlers.PatchTask)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	e.Start("localhost:8080")
}
