package db

import (
	"app/internal/TaskService"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatalf("Could not connection to db: %v", err)
	}
	DB.AutoMigrate(&TaskService.Task{})
	return DB, nil
}
