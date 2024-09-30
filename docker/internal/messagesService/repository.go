package messagesService

import "gorm.io/gorm"

type MessageRepository interface {
	// CreateMessage - Передаем в функцию message типа Message из orm.go
	// возвращаем созданный Message и ошибку
	CreateMessage(message Message) (Message, error)
	// GetAllMessages - Возвращаем массив из всех писем в БД и ошибку
	GetAllMessages() ([]Message, error)
	// UpdateMessageByID - Передаем id и Message, возвращаем обновленный Message
	// и ошибку
	UpdateMessageByID(id int, message Message) (Message, error)
	// DeleteMessageByID - Передаем id для удаления, возвращаем только ошибку
	DeleteMessageByID(id int) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

// (r *messageRepository) привязывает данную функцию к нашему репозиторию
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	result := r.db.Model(&message).Where("id = ?", id).Update("text", message.Text)

	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) DeleteMessageByID(id int) error {
	var message Message
	err := r.db.Where("id = ?", id).Delete(&message)
	if err != nil {
		return err.Error
	}
	return nil
}
