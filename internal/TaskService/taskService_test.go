package TaskService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name:  "успешное создание задачи",
			input: "Test task",
			mockSetup: func(m *MockTaskRepository) {
				m.On("CreateTask", mock.MatchedBy(func(task Task) bool {
					return task.Task == "Test task" && task.Status == "active"
				})).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при создании",
			input: "Bad task",
			mockSetup: func(m *MockTaskRepository) {
				m.On("CreateTask", mock.MatchedBy(func(task Task) bool {
					return task.Task == "Bad task" && task.Status == "active"
				})).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			result, err := service.CreateTask(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Task{}, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result.Task)
				assert.Equal(t, "active", result.Status)
				assert.NotEmpty(t, result.ID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetALLTasks(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное получение задач",
			mockSetup: func(m *MockTaskRepository) {
				expectedTasks := []Task{
					{ID: "1", Task: "Task 1", Status: "active"},
					{ID: "2", Task: "Task 2", Status: "completed"},
				}
				m.On("GetALLTasks").Return(expectedTasks, nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при получении",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetALLTasks").Return([]Task{}, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewTaskService(mockRepo)
			result, err := service.GetAllTasks()

			if tt.wantErr {
				assert.Error(t, err)
				assert.NotNil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Len(t, result, 2)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		mockSetup func(m *MockTaskRepository, id string)
		wantErr   bool
	}{
		{
			name:  "успешное получение задачи",
			input: "123",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("GetTaskById", id).Return(Task{
					ID:     id,
					Task:   "Test task",
					Status: "active",
				}, nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при получении",
			input: "999",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("GetTaskById", id).Return(Task{}, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.input)

			service := NewTaskService(mockRepo)
			result, err := service.GetTaskById(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Task{}, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result.ID)
				assert.Equal(t, "Test task", result.Task)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		taskText  string
		mockSetup func(m *MockTaskRepository, id, taskText string)
		wantErr   bool
	}{
		{
			name:     "успешное обновление задачи",
			id:       "123",
			taskText: "Updated task",
			mockSetup: func(m *MockTaskRepository, id, taskText string) {
				m.On("GetTaskById", id).Return(Task{
					ID:     id,
					Task:   "Old task",
					Status: "active",
				}, nil)
				m.On("UpdateTask", Task{
					ID:     id,
					Task:   taskText,
					Status: "active",
				}).Return(nil)
			},
			wantErr: false,
		},
		{
			name:     "ошибка при обновлении",
			id:       "999",
			taskText: "Updated task",
			mockSetup: func(m *MockTaskRepository, id, taskText string) {
				m.On("GetTaskById", id).Return(Task{}, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id, tt.taskText)

			service := NewTaskService(mockRepo)
			result, err := service.UpdateTask(tt.id, tt.taskText)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, Task{}, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.id, result.ID)
				assert.Equal(t, tt.taskText, result.Task)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		mockSetup func(m *MockTaskRepository, id string)
		wantErr   bool
	}{
		{
			name:  "успешное удаление задачи",
			input: "123",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteTask", id).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при удалении",
			input: "999",
			mockSetup: func(m *MockTaskRepository, id string) {
				m.On("DeleteTask", id).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.input)

			service := NewTaskService(mockRepo)
			err := service.DeleteTask(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
