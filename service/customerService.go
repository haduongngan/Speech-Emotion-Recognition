package service

import (
	"spser/model"
	"spser/repository"
)

type customerService struct {
	customerRepository model.CustomerRepository
}

type CustomerService interface {
	GetAll() ([]model.Customer, error)
	GetAllCall(userId int) ([]model.Call, error)
	FilterCallInTime(payload *model.CallTimeFilterPayload) ([]model.Call, error)
	CreateCustomer(call *model.Customer) error
	UpdateCustomer(call *model.Customer) error
	GetById(id int) (*model.Customer, error)
	DeleteCustomer(id int) error
}

func (s *customerService) GetAll() ([]model.Customer, error) {
	return s.customerRepository.GetAll()
}

func (s *customerService) CreateCustomer(call *model.Customer) error {
	return s.customerRepository.CreateCustomer(call)
}

func (s *customerService) UpdateCustomer(call *model.Customer) error {
	return s.customerRepository.UpdateCustomer(call)
}

func (s *customerService) GetById(id int) (*model.Customer, error) {
	return s.customerRepository.GetById(id)
}

func (s *customerService) DeleteCustomer(id int) error {
	return s.customerRepository.DeleteCustomer(id)
}

func (s *customerService) GetAllCall(userId int) ([]model.Call, error) {
	return s.customerRepository.GetAllCall(userId)
}

func (s *customerService) FilterCallInTime(payload *model.CallTimeFilterPayload) ([]model.Call, error) {
	return s.customerRepository.FilterCallInTime(payload)
}
func NewCustomerService() CustomerService {
	return &customerService{
		customerRepository: repository.NewCustomerRepository(),
	}
}
