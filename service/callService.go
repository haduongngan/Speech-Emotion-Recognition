package service

import (
	"spser/model"
	"spser/repository"
)

type callService struct {
	callRepository model.CallRepository
}

type CallService interface {
	GetAll() ([]model.Call, error)
	CreateCall(call *model.Call) error
	CreateMultiCall(calls []model.Call) error
	UpdateCall(call *model.Call) error
	GetById(id int) (*model.Call, error)
	DeleteCall(id int) error
}

func (s *callService) GetAll() ([]model.Call, error) {
	return s.callRepository.GetAll()
}

func (s *callService) CreateCall(call *model.Call) error {
	return s.callRepository.CreateCall(call)
}

func (s *callService) CreateMultiCall(calls []model.Call) error {
	return s.callRepository.CreateMultiCall(calls)
}
func (s *callService) UpdateCall(call *model.Call) error {
	return s.callRepository.UpdateCall(call)
}

func (s *callService) GetById(id int) (*model.Call, error) {
	return s.callRepository.GetById(id)
}

func (s *callService) DeleteCall(id int) error {
	return s.callRepository.DeleteCall(id)
}

func NewCallService() CallService {
	return &callService{
		callRepository: repository.NewCallRepository(),
	}
}
