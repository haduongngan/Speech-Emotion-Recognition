package repository

import (
	"spser/infrastructure"
	"spser/model"

	"gorm.io/gorm"
)

type callRepository struct{}

func (r *callRepository) GetAll() ([]model.Call, error) {
	db := infrastructure.GetDB()

	var calls []model.Call

	if err := db.Model(&model.Call{}).Find(&calls).Error; err != nil {
		return nil, err
	}

	return calls, nil
}

func (r *callRepository) CreateCall(call *model.Call) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Call{}).Create(call).Error; err != nil {
		return err
	}

	return nil
}

func (r *callRepository) UpdateCall(call *model.Call) error {
	db := infrastructure.GetDB()

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&model.Call{}).Updates(call).Error; err != nil {
		return err
	}

	return nil
}

func (r *callRepository) GetById(id int) (*model.Call, error) {
	db := infrastructure.GetDB()

	var call model.Call

	if err := db.Model(&model.Call{}).Where("id = ?", id).First(&call).Error; err != nil {
		return nil, err
	}

	return &call, nil
}
