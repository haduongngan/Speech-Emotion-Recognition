package service

import (
	"spser/model"
	"spser/repository"
)

type employeeService struct {
	employeeRepository model.EmployeeRepository
}

type EmployeeService interface {
	GetAll() ([]model.Employee, error)
	GetAllCall(userId int) ([]model.Call, error)
	FilterCallInTime(payload *model.CallTimeFilterPayload) ([]model.Call, error)
	CreateEmployee(call *model.Employee) error
	UpdateEmployee(call *model.Employee) error
	GetById(id int) (*model.Employee, error)
	DeleteEmployee(id int) error
}

func (s *employeeService) GetAll() ([]model.Employee, error) {
	return s.employeeRepository.GetAll()
}

func (s *employeeService) CreateEmployee(call *model.Employee) error {
	return s.employeeRepository.CreateEmployee(call)
}

func (s *employeeService) UpdateEmployee(call *model.Employee) error {
	return s.employeeRepository.UpdateEmployee(call)
}

func (s *employeeService) GetById(id int) (*model.Employee, error) {
	return s.employeeRepository.GetById(id)
}

func (s *employeeService) DeleteEmployee(id int) error {
	return s.employeeRepository.DeleteEmployee(id)
}

func (s *employeeService) GetAllCall(userId int) ([]model.Call, error) {
	return s.employeeRepository.GetAllCall(userId)
}

func (s *employeeService) FilterCallInTime(payload *model.CallTimeFilterPayload) ([]model.Call, error) {
	return s.employeeRepository.FilterCallInTime(payload)
}

func NewEmployeeService() EmployeeService {
	return &employeeService{
		employeeRepository: repository.NewEmployeeRepository(),
	}
}
