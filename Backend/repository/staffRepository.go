package repository

import (
	"spser/infrastructure"
	"spser/model"

	"gorm.io/gorm"
)

type staffRepository struct{}

func (r *staffRepository) GetAll() ([]model.Staff, error) {
	db := infrastructure.GetDB()

	var staffs []model.Staff

	if err := db.Model(&model.Staff{}).Find(&staffs).Error; err != nil {
		return nil, err
	}

	return staffs, nil
}

func (r *staffRepository) GetByUserId(userId int) (*model.Staff, error) {
	db := infrastructure.GetDB()

	var staff *model.Staff
	if err := db.Model(&model.Staff{}).Where("user_id = ?", userId).First(&staff).Error; err != nil {
		return nil, err
	}

	return staff, nil
}

func (r *staffRepository) GetAllCall(userId int) ([]model.Call, error) {
	db := infrastructure.GetDB()

	staff, err := r.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	var calls []model.Call
	if err := db.Preload("Segments").Model(&model.Call{}).Where("Staff_id = ?", staff.Id).Find(&calls).Error; err != nil {
		return nil, err
	}

	return calls, nil
}

func (r *staffRepository) FilterCallInTime(payload *model.StaffCallFilterPayload) ([]model.Call, error) {
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

func (r *staffRepository) CreateStaff(staff *model.Staff) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Staff{}).Create(staff).Error; err != nil {
		return err
	}

	return nil

}

func (r *staffRepository) UpdateStaff(staffName *model.StaffNameUpdate) error {
	db := infrastructure.GetDB()

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&model.Staff{}).Where("id = ?", staffName.Id).Update("name", staffName.Name).Error; err != nil {
		return err
	}

	return nil
}

func (r *staffRepository) GetById(id int) (*model.Staff, error) {
	db := infrastructure.GetDB()

	var staff *model.Staff
	if err := db.Model(&model.Staff{}).Where("id = ?", id).First(&staff).Error; err != nil {
		return nil, err
	}

	return staff, nil
}

func (r *staffRepository) DeleteStaff(id int) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Staff{}).Where("id = ?", id).Delete(&model.Staff{}).Error; err != nil {
		return err
	}

	return nil
}

func NewStaffRepository() model.StaffRepository {
	return &staffRepository{}
}
