package UserService

import (
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(email, password string) (User, error)
	GetAllUsers() ([]User, error)
	GetUserById(id string) (User, error)
	GetUserByEmail(email string) (User, error)
	UpdateUser(id, Email string) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (r *userService) CreateUser(email, password string) (User, error) {
	newUser := User{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := r.repo.CreateUser(newUser); err != nil {
		return User{}, err
	}
	return newUser, nil
}

func (r *userService) GetAllUsers() ([]User, error) {
	return r.repo.GetAllUsers()
}

func (r *userService) GetUserById(id string) (User, error) {
	return r.repo.GetUserByID(id)
}

func (r *userService) GetUserByEmail(email string) (User, error) {
	return r.repo.GetUserByEmail(email)
}

func (r *userService) UpdateUser(id, Email string) (User, error) {
	user, err := r.GetUserById(id)
	if err != nil {
		return User{}, err
	}
	user.Email = Email
	user.UpdatedAt = time.Now()
	if err := r.repo.UpdateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *userService) DeleteUser(id string) error {
	return r.repo.DeleteUser(id)
}
