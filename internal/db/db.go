package db

import (
	"app/internal/TaskService"
	"app/internal/UserService"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connection to db:%v", err)
	}
	if err := DB.AutoMigrate(&TaskService.Task{}); err != nil {
		log.Fatalf("Could not migrate db:%v", err)
	}
	if err := DB.AutoMigrate(&UserService.User{}); err != nil {
		log.Fatalf("Could not migrate db:%v", err)
	}
	return DB, nil

}
