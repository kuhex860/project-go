package handlers

import (
	"app/internal/TaskService"
	"app/internal/web/tasks"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type TaskHandler struct {
	service TaskService.TaskService
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, task := range allTasks {
		task := tasks.Task{
			Id:     task.ID,
			Status: task.Status,
			Task:   task.Task,
			UserId: task.UserID.String(),
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	body := request.Body
	if body.UserId == "" {
		return nil, fmt.Errorf("userId is required")
	}
	userID, err := uuid.Parse(body.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid userId format: %v", err)
	}
	createdTask, err := h.service.CreateTask(body.Task, userID)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     createdTask.ID,
		Task:   createdTask.Task,
		Status: createdTask.Status,
		UserId: createdTask.UserID.String(),
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id
	err := h.service.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			return tasks.DeleteTasksId404Response{}, nil
		}
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	body := request.Body
	updatedTask, err := h.service.UpdateTask(id, *body.Task)
	if err != nil {
		if err.Error() == "task not found" {
			return tasks.PatchTasksId404Response{}, nil
		}
		return nil, err

	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     updatedTask.ID,
		Status: updatedTask.Status,
		Task:   updatedTask.Task,
		UserId: updatedTask.UserID.String(),
	}
	return response, nil

}

func NewTaskHandler(s TaskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}
