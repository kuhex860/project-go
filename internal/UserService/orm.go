package UserService

import "time"

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"update_at"`
	CreatedAt time.Time `json:"create_at"`
}

type UserUpdate struct {
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"update_at"`
}
