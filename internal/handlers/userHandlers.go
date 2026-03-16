package handlers

import (
	"app/internal/UserService"
	"app/internal/web/users"
	"context"
)

type UserHandler struct {
	service UserService.UserService
}

func (u UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		user := users.User{
			Id:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		}
		response = append(response, user)
	}
	return response, nil
}

func (u UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	body := request.Body
	createdUser, err := u.service.CreateUser(body.Email, body.Password)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:        createdUser.ID,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
		DeletedAt: createdUser.DeletedAt,
	}
	return response, nil
}

func (u UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := u.service.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			return users.DeleteUsersId404Response{}, nil
		}
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (u UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	body := request.Body
	updatedUser, err := u.service.UpdateUser(id, *body.Email)
	if err != nil {
		if err.Error() == "user not found" {
			return users.PatchUsersId404Response{}, nil
		}
		return nil, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:        updatedUser.ID,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
		DeletedAt: updatedUser.DeletedAt,
	}
	return response, nil
}

func (u UserHandler) GetUsersIdTasks(ctx context.Context, request users.GetUsersIdTasksRequestObject) (users.GetUsersIdTasksResponseObject, error) {
	userID := request.Id
	tasks, err := u.service.GetTasksForUser(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return users.GetUsersIdTasks404Response{}, nil
		}
		return nil, err
	}
	response := users.GetUsersIdTasks200JSONResponse{}
	for _, task := range tasks {
		response = append(response, users.Task{
			Id:     task.ID,
			Task:   task.Task,
			Status: task.Status,
			UserId: task.UserID.String(),
		})
	}
	return response, nil
}

func NewUserHandler(s UserService.UserService) *UserHandler {
	return &UserHandler{service: s}
}
