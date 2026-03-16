package TaskService

import (
	"app/internal/models"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task string, userID uuid.UUID) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id string) (models.Task, error)
	GetTaskByUserID(userID uuid.UUID) ([]models.Task, error)
	UpdateTask(id, task string) (models.Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}
func (r *taskService) CreateTask(task string, userID uuid.UUID) (models.Task, error) {
	newTask := models.Task{
		ID:     uuid.NewString(),
		Task:   task,
		Status: "active",
		UserID: userID,
	}
	if err := r.repo.CreateTask(newTask); err != nil {
		return models.Task{}, err
	}
	return newTask, nil
}

func (r *taskService) GetAllTasks() ([]models.Task, error) {
	return r.repo.GetAllTasks()
}

func (r *taskService) GetTaskById(id string) (models.Task, error) {
	return r.repo.GetTaskById(id)
}
func (r *taskService) GetTaskByUserID(userID uuid.UUID) ([]models.Task, error) {
	return r.repo.GetTaskByUserID(userID)
}

func (r *taskService) UpdateTask(id, task string) (models.Task, error) {
	updateTask, err := r.repo.GetTaskById(id)
	if err != nil {
		return models.Task{}, err
	}
	updateTask.Task = task
	if err := r.repo.UpdateTask(updateTask); err != nil {
		return models.Task{}, err
	}
	return updateTask, nil
}
func (r *taskService) DeleteTask(id string) error {
	return r.repo.DeleteTask(id)
}
