package userService

import "gorm.io/gorm"

type UserRepository interface {
	// Create user - Передаем в функцию user типа User из orm.go
	// возвращаем созданный User и ошибку
	CreateUser(user User) (User, error)
	// GetAllUsers- Возвращаем массив из всех User в БД и ошибку
	GetAllUsers() ([]User, error)
	// UpdateUserByID - Передаем id и User, возвращаем обновленный User
	// и ошибку
	UpdateUserByID(id int, user User) (User, error)
	// DeleteUserByID - Передаем id для удаления, возвращаем только ошибку
	DeleteUserByID(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// (r *messageRepository) привязывает данную функцию к нашему репозиторию
func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id int, user User) (User, error) {
	result := r.db.Model(&user).Where("id = ?", id).Update("password", user.Password)

	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(id int) error {
	var user User
	err := r.db.Where("id = ?", id).Delete(&user)
	if err != nil {
		return err.Error
	}
	return nil
}
