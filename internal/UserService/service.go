package UserService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(email, password string) (User, error)
	GetALLUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id string, update UserUpdate) (User, error)
	DeleteUser(id string) error
	VerifyPassword(user User, password string) bool
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(email, password string) (User, error) {
	existingUser, err := s.repo.GetUserByEmail(email)
	if err == nil && existingUser.ID != "" {
		return User{}, errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) GetALLUsers() ([]User, error) {
	return s.repo.GetALLUsers()
}

func (s *userService) GetUserByID(id string) (User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) UpdateUser(id string, update UserUpdate) (User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, errors.New("user not found")
	}
	if update.Email != nil {
		existingUser, err := s.repo.GetUserByEmail(*update.Email)
		if err == nil && existingUser.ID != "" && existingUser.ID != id {
			return User{}, errors.New("email already in use")
		}
		user.Email = *update.Email
	}
	if update.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*update.Password), bcrypt.DefaultCost)
		if err != nil {
			return User{}, err
		}
		user.Password = string(hashedPassword)
	}
	user.UpdatedAt = time.Now()

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.SoftDeleteUser(id)
}

func (s *userService) VerifyPassword(user User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
