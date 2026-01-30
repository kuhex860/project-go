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

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Status: &tsk.Status,
			Task:   &tsk.Task,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := TaskService.Task{
		Task: *taskRequest.Task,
	}
	createdTask, err := h.service.CreateTask(taskToCreate.Task)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		Status: &createdTask.Status,
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
	taskRequest := request.Body
	taskToUpdate := TaskService.Task{
		Task: *taskRequest.Task,
	}
	updatedTask, err := h.service.UpdateTask(id, taskToUpdate.Task)
	if err != nil {
		if err.Error() == "task not found" {
			return tasks.PatchTasksId404Response{}, nil
		}
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		Status: &updatedTask.Status,
	}
	return response, nil
}

func NewTaskHandler(s TaskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

//func (h *TaskHandler) GetTasks(c echo.Context) error {
//	tasks, err := h.service.GetAllTasks()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
//	}
//	return c.JSON(http.StatusOK, tasks)
//}
//
//func (h *TaskHandler) PostTask(c echo.Context) error {
//	var req TaskService.TaskRequest
//	if err := c.Bind(&req); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
//	}
//
//	newTasks, err := h.service.CreateTask(req.Task)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create task"})
//	}
//
//	return c.JSON(http.StatusCreated, newTasks)
//}
//
//func (h *TaskHandler) PatchTask(c echo.Context) error {
//	id := c.Param("id")
//	var req TaskService.TaskRequest
//	if err := c.Bind(&req); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
//	}
//
//	updatedTask, err := h.service.UpdateTask(id, req.Task)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
//	}
//	return c.JSON(http.StatusOK, updatedTask)
//}
//
//func (h *TaskHandler) DeleteTask(c echo.Context) error {
//	id := c.Param("id")
//
//	if err := h.service.DeleteTask(id); err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
//	}
//	return c.NoContent(http.StatusNoContent)
//}
