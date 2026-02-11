package UserService

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) error
	GetALLUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	GetUserByEmail(email string) (User, error)
	UpdateUser(user User) error
	DeleteUser(id string) error
	SoftDeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetALLUsers() ([]User, error) {
	var users []User
	err := r.db.Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id string) (User, error) {
	var user User
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return user, err
}

func (r *userRepository) GetUserByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(user User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&User{}).Error
}

func (r *userRepository) SoftDeleteUser(id string) error {
	now := time.Now()
	return r.db.Model(&User{}).Where("id = ?", id).Update("deleted_at", now).Error
}
