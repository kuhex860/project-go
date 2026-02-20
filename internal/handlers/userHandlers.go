package handlers

import (
	"app/internal/UserService"
	"app/internal/web/users"
	"context"
	"time"
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
		}
		response = append(response, user)
	}
	return response, nil
}

func (u UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	user := request.Body
	Create := UserService.User{
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createdUser, err := u.service.CreateUser(Create.Email, Create.Password)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:        createdUser.ID,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
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
	}
	return users.DeleteUsersId204Response{}, nil
}

func (u UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	user := request.Body
	update := UserService.User{
		Email: *user.Email,
	}
	updatedUser, err := u.service.UpdateUser(id, update.Email)
	if err != nil {
		if err.Error() == "user not found" {
			return users.PatchUsersId404Response{}, nil
		}
	}
	response := users.PatchUsersId200JSONResponse{
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}
	return response, nil
}

func NewUserHandler(s UserService.UserService) *UserHandler {
	return &UserHandler{service: s}
}
