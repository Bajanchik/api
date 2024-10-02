package userService

type UserService struct {
	UserRepo UserRepository
}

func NewService(UserRepo UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id int, user User) (User, error) {
	return s.UserRepo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id int) error {
	return s.UserRepo.DeleteUserByID(id)
}
