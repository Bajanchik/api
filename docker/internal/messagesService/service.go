package messagesService

type MessageService struct {
	messageRepo MessageRepository
}

func NewService(messageRepo MessageRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo}
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.messageRepo.CreateMessage(message)
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.messageRepo.GetAllMessages()
}

func (s *MessageService) UpdateMessageByID(id int, message Message) (Message, error) {
	return s.messageRepo.UpdateMessageByID(id, message)
}

func (s *MessageService) DeleteMessageByID(id int) error {
	return s.messageRepo.DeleteMessageByID(id)
}
