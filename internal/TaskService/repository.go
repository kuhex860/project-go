package TaskService

import (
	"app/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id string) (models.Task, error)
	GetTaskByUserID(userID uuid.UUID) ([]models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task models.Task) error {
	return r.db.Create(&task).Error
}
func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *taskRepository) GetTaskById(id string) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) GetTaskByUserID(userID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
func (r *taskRepository) UpdateTask(task models.Task) error {
	return r.db.Save(&task).Error
}
func (r *taskRepository) DeleteTask(id string) error {
	return r.db.Delete(&models.Task{}, "id = ?", id).Error
}
