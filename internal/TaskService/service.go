package TaskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(task string) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTask(id, task string) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}
func (r *taskService) CreateTask(task string) (Task, error) {
	newTask := Task{
		ID:     uuid.NewString(),
		Task:   task,
		Status: "active",
	}
	if err := r.repo.CreateTask(newTask); err != nil {
		return Task{}, err
	}
	return newTask, nil
}

func (r *taskService) GetAllTasks() ([]Task, error) {
	return r.repo.GetAllTasks()
}

func (r *taskService) GetTaskById(id string) (Task, error) {
	return r.repo.GetTaskById(id)
}

func (r *taskService) UpdateTask(id, task string) (Task, error) {
	updateTask, err := r.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}
	updateTask.Task = task
	if err := r.repo.UpdateTask(updateTask); err != nil {
		return Task{}, err
	}
	return updateTask, nil
}
func (r *taskService) DeleteTask(id string) error {
	return r.repo.DeleteTask(id)
}
