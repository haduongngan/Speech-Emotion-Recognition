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

func (r *employeeRepository) GetByUserId(userId int) (*model.Employee, error) {
	db := infrastructure.GetDB()

	var employee *model.Employee
	if err := db.Model(&model.Employee{}).Where("user_id = ?", userId).First(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) GetAllCall(userId int) ([]model.Call, error) {
	db := infrastructure.GetDB()

	employee, err := r.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	var calls []model.Call
	if err := db.Preload("Segments").Model(&model.Call{}).Where("employee_id = ?", employee.Id).Find(&calls).Error; err != nil {
		return nil, err
	}

	return calls, nil
}

func (r *employeeRepository) FilterCallInTime(payload *model.CallTimeFilterPayload) ([]model.Call, error) {
	var callsInTime []model.Call
	allCallsOfUser, err := r.GetAllCall(payload.UserId)
	if err != nil {
		return nil, err
	}

	for i := range allCallsOfUser {
		if allCallsOfUser[i].StartTime.After(payload.StartTime) && allCallsOfUser[i].StartTime.Before(payload.EndTime) {
			callsInTime = append(callsInTime, allCallsOfUser[i])
		}
	}

	return callsInTime, nil
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

func (r *employeeRepository) GetById(id int) (*model.Employee, error) {
	db := infrastructure.GetDB()

	var employee *model.Employee
	if err := db.Model(&model.Employee{}).Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) DeleteEmployee(id int) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Employee{}).Where("id = ?", id).Delete(&model.Employee{}).Error; err != nil {
		return err
	}

	return nil
}

func NewEmployeeRepository() model.EmployeeRepository {
	return &employeeRepository{}
}
