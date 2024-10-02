package handlers

import (
	"context"
	"docker/internal/messagesService" // Импортируем наш сервис
	"docker/internal/web/messages"
	"errors"

	"gorm.io/gorm"
)

type MessageHandler struct {
	Service *messagesService.MessageService
}

// DeleteMessages implements messages.StrictServerInterface.
func (h *MessageHandler) DeleteMessages(ctx context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
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
func (h *MessageHandler) PatchMessages(ctx context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
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
func (h *MessageHandler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
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

func (h *MessageHandler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body

	// Проверка на nil для messageRequest.Message
	if messageRequest.Message == nil {
		return messages.PostMessages201JSONResponse{}, errors.New("message cannot be nil")
	}

	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}

	// Проверка на nil для h.Service
	if h.Service == nil {
		return messages.PostMessages201JSONResponse{}, errors.New("service is not initialized")
	}

	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return messages.PostMessages201JSONResponse{}, err
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

func NewMessageHandler(service *messagesService.MessageService) *MessageHandler {
	return &MessageHandler{
		Service: service,
	}
}
