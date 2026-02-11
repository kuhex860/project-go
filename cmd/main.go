package main

import (
	"app/internal/TaskService"
	"app/internal/UserService"
	"app/internal/db"
	"app/internal/handlers"
	"app/internal/web/tasks"
	"app/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.DB.AutoMigrate(&TaskService.Task{}); err != nil {
		log.Printf("Warning: Failed to auto-migrate tasks table: %v", err)

	}

	// Автомиграция для пользователей
	if err := db.DB.AutoMigrate(&UserService.User{}); err != nil {
		log.Printf("Warning: Failed to auto-migrate users table: %v", err)

	}

	tasksRepo := TaskService.NewTaskRepository(db.DB)
	tasksService := TaskService.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	usersRepo := UserService.NewUserRepository(db.DB)
	usersService := UserService.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
	}))

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	log.Println("Server starting on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
