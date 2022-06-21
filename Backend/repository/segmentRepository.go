package repository

import (
	"spser/infrastructure"
	"spser/model"
)

type segmentRepository struct{}

func (r *segmentRepository) GetAll() ([]model.Segment, error) {
	db := infrastructure.GetDB()

	var segments []model.Segment

	if err := db.Model(&model.Segment{}).Find(&segments).Error; err != nil {
		return nil, err
	}

	return segments, nil
}

func (r *segmentRepository) GetById(id int) (*model.Segment, error) {
	db := infrastructure.GetDB()

	var segment model.Segment

	if err := db.Model(&model.Segment{}).Where("id = ?", id).First(&segment).Error; err != nil {
		return nil, err
	}

	return &segment, nil
}

func (r *segmentRepository) CreateSegment(segment *model.Segment) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Segment{}).Create(&segment).Error; err != nil {
		return err
	}

	return nil
}

func (r *segmentRepository) CreateMultiSegment(segment []model.Segment) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Segment{}).Create(&segment).Error; err != nil {
		return err
	}

	return nil
}

func (r *segmentRepository) GetByCallId(callId int) ([]model.Segment, error) {
	db := infrastructure.GetDB()

	var segments []model.Segment

	if err := db.Model(&model.Segment{}).Where("call_id = ?", callId).Find(&segments).Error; err != nil {
		return nil, err
	}

	return segments, nil
}

func (r *segmentRepository) DeleteSegment(id int) error {
	db := infrastructure.GetDB()

	if err := db.Delete(&model.Segment{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *segmentRepository) GetEmotion(id int) (string, error) {
	db := infrastructure.GetDB()

	var emotion string

	if err := db.Model(&model.Segment{}).Where("id = ?", id).Select("emotion").First(&emotion).Error; err != nil {
		return "get emo error", err
	}

	return emotion, nil
}

func NewSegmentRepository() model.SegmentRepository {
	return &segmentRepository{}
}
