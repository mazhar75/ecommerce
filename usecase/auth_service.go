package usecase

import "github/ecommerce/domain/user"

type AuthService struct {
	Repo user.UserRepository
}

func NewAuthService(repo user.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}
func (s *AuthService) CreateUser(u user.Users) error {
	return s.Repo.CreateUser(u)
}
func (s *AuthService) Login(email string, password string) error {
	return s.Repo.Login(email, password)
}
func (s *AuthService) DeleteUser(uId int) error {
	return s.Repo.DeleteUser(uId)
}
