package service

import (
	"spser/model"
	"spser/repository"
)

type staffService struct {
	staffRepository model.StaffRepository
}

type StaffService interface {
	GetAll() ([]model.Staff, error)
	GetAllCall(userId int) ([]model.Call, error)
	FilterCallInTime(payload *model.StaffCallFilterPayload) ([]model.Call, error)
	CreateStaff(staff *model.Staff) error
	UpdateStaff(staffName *model.StaffNameUpdate) error
	GetById(id int) (*model.Staff, error)
	DeleteStaff(id int) error
}

func (s *staffService) GetAll() ([]model.Staff, error) {
	return s.staffRepository.GetAll()
}

func (s *staffService) CreateStaff(call *model.Staff) error {
	return s.staffRepository.CreateStaff(call)
}

func (s *staffService) UpdateStaff(staffName *model.StaffNameUpdate) error {
	return s.staffRepository.UpdateStaff(staffName)
}

func (s *staffService) GetById(id int) (*model.Staff, error) {
	return s.staffRepository.GetById(id)
}

func (s *staffService) DeleteStaff(id int) error {
	return s.staffRepository.DeleteStaff(id)
}

func (s *staffService) GetAllCall(userId int) ([]model.Call, error) {
	return s.staffRepository.GetAllCall(userId)
}

func (s *staffService) FilterCallInTime(payload *model.StaffCallFilterPayload) ([]model.Call, error) {
	return s.staffRepository.FilterCallInTime(payload)
}

func NewStaffService() StaffService {
	return &staffService{
		staffRepository: repository.NewStaffRepository(),
	}
}
