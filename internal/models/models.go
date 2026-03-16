package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID     string    `gorm:"primaryKey" json:"id"`
	Task   string    `json:"task"`
	Status string    `json:"status"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}

type User struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Tasks     []Task     `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}
