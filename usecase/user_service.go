package usecase

import "github/ecommerce/domain/user"

type UserService struct {
	Repo user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{Repo: repo}
}
func (s *UserService) CreateUser(u user.Users) error {
	return s.Repo.CreateUser(u)
}
func (s *UserService) Login(email string, password string) error {
	return s.Repo.Login(email, password)
}
func (s *UserService) DeleteUser(uId int) error {
	return s.Repo.DeleteUser(uId)
}
