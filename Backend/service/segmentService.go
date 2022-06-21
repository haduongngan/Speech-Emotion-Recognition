package service

import (
	"spser/model"
	"spser/repository"
)

type segmentService struct {
	segmentRepository model.SegmentRepository
}

type SegmentService interface {
	GetAll() ([]model.Segment, error)
	GetById(id int) (*model.Segment, error)
	GetByCallId(callId int) ([]model.Segment, error)
	CreateSegment(new *model.Segment) error
	CreateMultiSegment(segments []model.Segment) error
	DeleteSegment(id int) error
	GetEmotion(id int) (string, error)
}

func (s *segmentService) GetAll() ([]model.Segment, error) {
	return s.segmentRepository.GetAll()
}

func (s *segmentService) GetById(id int) (*model.Segment, error) {
	return s.segmentRepository.GetById(id)
}

func (s *segmentService) GetByCallId(callId int) ([]model.Segment, error) {
	return s.segmentRepository.GetByCallId(callId)
}

func (s *segmentService) CreateSegment(new *model.Segment) error {
	return s.segmentRepository.CreateSegment(new)
}

func (s *segmentService) CreateMultiSegment(segments []model.Segment) error {
	return s.segmentRepository.CreateMultiSegment(segments)
}

func (s *segmentService) DeleteSegment(id int) error {
	return s.segmentRepository.DeleteSegment(id)
}

func (s *segmentService) GetEmotion(id int) (string, error) {
	return s.segmentRepository.GetEmotion(id)
}

func NewSegmentService() SegmentService {
	return &segmentService{
		segmentRepository: repository.NewSegmentRepository(),
	}
}
