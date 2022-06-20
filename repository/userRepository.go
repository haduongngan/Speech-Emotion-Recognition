package repository

import (
	"spser/infrastructure"
	"spser/model"
	"time"
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

func (r *userRepository) CreateUser(newUser *model.User) (*model.User, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&model.User{}).Create(&newUser).Error; err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *userRepository) GetById(id int) (*model.User, error) {
	db := infrastructure.GetDB()

	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	db := infrastructure.GetDB()

	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) DeleteUser(id int) (*model.User, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&model.User{Id: id}).Update("deletedAt", time.Now()).Error; err != nil {
		return nil, err
	}

	return r.GetById(id)
}

func (r *userRepository) LoginTokenRequest(user *model.User) (bool, error) {
	db := infrastructure.GetDB()

	var userInfo model.User
	if err := db.Where(&model.User{
		Username: user.Username,
		Password: user.Password,
	}).First(&userInfo).Error; err != nil {
		infrastructure.ErrLog.Println(err)
		return false, nil
	}

	user.ExpiresAt = time.Now().Local().Add(time.Hour*time.Duration(infrastructure.GetExtendAccessHour())).UnixNano() / 100000000
	return true, nil
}
func NewUserRepository() model.UserRepository {
	return &userRepository{}
}
