package handlers

import (
	"app/internal/UserService"
	"app/internal/web/users"
	"context"
)

type UserHandler struct {
	service UserService.UserService
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetALLUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:        usr.ID,
			Email:     usr.Email,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
		}

		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	if userRequest == nil || userRequest.Email == "" || userRequest.Password == "" {
		return users.PostUsers400Response{}, nil
	}
	createdUser, err := h.service.CreateUser(userRequest.Email, userRequest.Password)
	if err != nil {
		if err.Error() == "user with this email already exists" {
			return users.PostUsers400Response{}, nil
		}
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

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := h.service.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			return users.DeleteUsersId404Response{}, nil
		}
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	userRequest := request.Body
	if userRequest == nil || (userRequest.Email == nil && userRequest.Password == nil) {
		return users.PatchUsersId400Response{}, nil
	}
	update := UserService.UserUpdate{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	updatedUser, err := h.service.UpdateUser(id, update)
	if err != nil {
		if err.Error() == "user not found" {
			return users.PatchUsersId404Response{}, nil
		}
		if err.Error() == "email already in use" {
			return users.PatchUsersId400Response{}, nil
		}
		return nil, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:        updatedUser.ID,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}
	return response, nil
}

func NewUserHandler(s UserService.UserService) *UserHandler {
	return &UserHandler{service: s}
}
