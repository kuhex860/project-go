package UserService

import (
	"app/internal/TaskService"
	"app/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(email, password string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(id, Email string) (models.User, error)
	DeleteUser(id string) error
	GetTasksForUser(userId string) ([]models.Task, error)
}

type userService struct {
	repo        UserRepository
	taskService TaskService.TaskService
}

func NewUserService(r UserRepository, ts TaskService.TaskService) UserService {
	return &userService{repo: r, taskService: ts}
}

func (r *userService) CreateUser(email, password string) (models.User, error) {
	newUser := models.User{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  password,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		DeletedAt: nil,
		Tasks:     nil,
	}
	if err := r.repo.CreateUser(newUser); err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (r *userService) GetAllUsers() ([]models.User, error) {
	return r.repo.GetAllUsers()
}

func (r *userService) GetUserById(id string) (models.User, error) {
	return r.repo.GetUserByID(id)
}

func (r *userService) GetUserByEmail(email string) (models.User, error) {
	return r.repo.GetUserByEmail(email)
}

func (r *userService) UpdateUser(id, Email string) (models.User, error) {
	user, err := r.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	user.Email = Email
	user.UpdatedAt = time.Now()
	if err := r.repo.UpdateUser(user); err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *userService) GetTasksForUser(userId string) ([]models.Task, error) {
	_, err := r.repo.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	return r.taskService.GetTaskByUserID(uid)
}

func (r *userService) DeleteUser(id string) error {
	return r.repo.DeleteUser(id)
}
