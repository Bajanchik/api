package handlers

import (
	"context"
	"docker/internal/userService" // Импортируем наш сервис
	"docker/internal/web/users"
	"errors"

	"gorm.io/gorm"
)

type UserHandler struct {
	Service *userService.UserService ///////!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
}

// DeleteMessages implements messages.StrictServerInterface.
func (h *UserHandler) DeleteUsers(ctx context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	userRequest := request.Body
	userToDelete := userService.User{Model: gorm.Model{ID: *userRequest.Id}}
	err := h.Service.DeleteUserByID(int(userToDelete.ID))

	if err != nil {
		return nil, err
	}

	response := users.DeleteUsers200JSONResponse{
		Id: &userToDelete.ID,
	}
	return response, nil

}

// PatchMessages implements messages.StrictServerInterface.
func (h *UserHandler) PatchUsers(ctx context.Context, request users.PatchUsersRequestObject) (users.PatchUsersResponseObject, error) {
	userRequest := request.Body
	userToUpdate := userService.User{Model: gorm.Model{ID: *userRequest.Id}, Password: *userRequest.Password}
	updatedUser, err := h.Service.UpdateUserByID(int(userToUpdate.ID), userToUpdate)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsers200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil

}

// GetMessages implements messages.StrictServerInterface.
func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	// Получение всех сообщений из сервиса
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := users.GetUsers200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body

	// Проверка на nil для messageRequest.Message
	if userRequest.Email == nil || userRequest.Password == nil {
		return users.PostUsers201JSONResponse{}, errors.New("user cannot be nil")
	}

	// Обращаемся к сервису и создаем сообщение
	userToCreate := userService.User{Email: *userRequest.Email, Password: *userRequest.Password}

	// Проверка на nil для h.Service
	if h.Service == nil {
		return users.PostUsers201JSONResponse{}, errors.New("service is not initialized")
	}

	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return users.PostUsers201JSONResponse{}, err
	}

	// создаем структуру респонс
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	// Просто возвращаем респонс!
	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
