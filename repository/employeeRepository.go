package repository

import (
	"spser/infrastructure"
	"spser/model"

	"gorm.io/gorm"
)

type employeeRepository struct{}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	db := infrastructure.GetDB()

	var employees []model.Employee

	if err := db.Model(&model.Employee{}).Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *employeeRepository) CreateEmployee(employee *model.Employee) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Employee{}).Create(employee).Error; err != nil {
		return err
	}

	return nil

}

func (r *employeeRepository) UpdateEmployee(employee *model.Employee) error {
	db := infrastructure.GetDB()

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&model.Employee{}).Updates(employee).Error; err != nil {
		return err
	}

	return nil
}

func NewEmployeeRepository() model.EmployeeRepository {
	return &employeeRepository{}
}
