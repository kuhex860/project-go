package main

import (
	"app/internal/TaskService"
	"app/internal/db"
	"app/internal/handlers"
	"app/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//	func main() {
//		database, err := db.InitDB()
//		if err != nil {
//			log.Fatalf("Could not connect to db: %v", err)
//		}
//		taskRepo := TaskService.NewTaskRepository(database)
//		taskService := TaskService.NewTaskService(taskRepo)
//		taskHandlers := handlers.NewTaskHandler(taskService)
//		e := echo.New()
//
//		e.Use(middleware.Logger())
//		e.Use(middleware.CORS())
//
//		e.GET("/tasks", taskHandlers.GetTasks)
//		e.POST("/tasks", taskHandlers.PostTask)
//		e.PATCH("/tasks/:id", taskHandlers.PatchTask)
//		e.DELETE("/tasks/:id", taskHandlers.DeleteTask)
//
//		e.Start("localhost:8080")
//	}
func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := db.DB.AutoMigrate(&TaskService.Task{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	repo := TaskService.NewTaskRepository(db.DB)
	service := TaskService.NewTaskService(repo)

	handler := handlers.NewTaskHandler(service)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
