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

func (t *taskService) CreateTask(task string) (Task, error) {
	newTasks := Task{
		ID:     uuid.NewString(),
		Task:   task,
		Status: "active",
	}
	if err := t.repo.CreateTask(newTasks); err != nil {
		return Task{}, err
	}
	return newTasks, nil
}

func (t *taskService) GetAllTasks() ([]Task, error) {
	return t.repo.GetALLTasks()
}

func (t *taskService) GetTaskById(id string) (Task, error) {
	return t.repo.GetTaskById(id)
}

func (t *taskService) UpdateTask(id, task string) (Task, error) {
	newTasks, err := t.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}
	newTasks.Task = task

	if err := t.repo.UpdateTask(newTasks); err != nil {
		return Task{}, err
	}
	return newTasks, nil
}

func (t *taskService) DeleteTask(id string) error {
	return t.repo.DeleteTask(id)
}
