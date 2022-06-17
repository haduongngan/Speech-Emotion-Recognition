package service

import (
	"spser/model"
	"spser/repository"
)

type userService struct {
	userRepository model.UserRepository
}

type UserService interface {
	GetAll() ([]model.User, error)
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.userRepository.GetAll()
}

func NewUserService() UserService {
	return &userService{
		userRepository: repository.NewUserRepository(),
	}
}
