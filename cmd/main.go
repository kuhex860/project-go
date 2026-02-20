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
	db.InitDB()
	db.DB.AutoMigrate(&TaskService.Task{})
	db.DB.AutoMigrate(&UserService.User{})

	repo := TaskService.NewTaskRepository(db.DB)
	service := TaskService.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	userRepo := UserService.NewUserRepository(db.DB)
	userService := UserService.NewUserService(userRepo)
	userhandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)
	strictHandlers := users.NewStrictHandler(userhandler, nil)
	users.RegisterHandlers(e, strictHandlers)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
