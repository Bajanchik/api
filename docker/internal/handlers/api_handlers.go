package handlers

import (
	"context"
	"docker/internal/messagesService" // Импортируем наш сервис
	"docker/internal/web/messages"

	"gorm.io/gorm"
)

type Handler struct {
	Service *messagesService.MessageService
}

// DeleteMessages implements messages.StrictServerInterface.
func (h *Handler) DeleteMessages(ctx context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToDelete := messagesService.Message{Model: gorm.Model{ID: *messageRequest.Id}}
	err := h.Service.DeleteMessageByID(int(messageToDelete.ID))

	if err != nil {
		return nil, err
	}

	response := messages.DeleteMessages201JSONResponse{
		Id: &messageToDelete.ID,
	}
	return response, nil

}

// PatchMessages implements messages.StrictServerInterface.
func (h *Handler) PatchMessages(ctx context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToUpdate := messagesService.Message{Model: gorm.Model{ID: *messageRequest.Id}, Text: *messageRequest.Message}
	updatedMessage, err := h.Service.UpdateMessageByID(int(messageToUpdate.ID), messageToUpdate)

	if err != nil {
		return nil, err
	}

	response := messages.PatchMessages201JSONResponse{
		Id:      &updatedMessage.ID,
		Message: &updatedMessage.Text,
	}

	return response, nil

}

// GetMessages implements messages.StrictServerInterface.
func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}
