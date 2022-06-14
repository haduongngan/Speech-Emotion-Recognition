package repository

import (
	"spser/infrastructure"
	"spser/model"
)

type userRepository struct{}

func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	db := infrastructure.GetDB()
	if err := db.Model(&model.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserRepository() model.UserRepository {
	return &userRepository{}
}
