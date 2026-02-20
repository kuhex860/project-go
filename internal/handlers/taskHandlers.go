package handlers

import (
	"app/internal/TaskService"
	"app/internal/web/tasks"
	"context"
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
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	task := request.Body
	Create := TaskService.Task{
		Task: task.Task,
	}
	createdTask, err := h.service.CreateTask(Create.Task)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     createdTask.ID,
		Status: createdTask.Status,
		Task:   createdTask.Task,
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
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	task := request.Body
	Update := TaskService.Task{
		Task: *task.Task,
	}
	updatedTask, err := h.service.UpdateTask(id, Update.Task)
	if err != nil {
		if err.Error() == "task not found" {
			return tasks.PatchTasksId404Response{}, nil
		}

	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     updatedTask.ID,
		Status: updatedTask.Status,
		Task:   updatedTask.Task,
	}
	return response, nil

}

func NewTaskHandler(s TaskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}
