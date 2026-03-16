package UserService

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id string) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *userRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}
func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}
func (r *userRepository) UpdateUser(user models.User) error {
	return r.db.Save(&user).Error
}
func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
