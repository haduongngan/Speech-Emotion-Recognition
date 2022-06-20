package repository

import (
	"spser/infrastructure"
	"spser/model"

	"gorm.io/gorm"
)

type customerRepository struct{}

func (r *customerRepository) GetAll() ([]model.Customer, error) {
	db := infrastructure.GetDB()

	var customers []model.Customer

	if err := db.Model(&model.Customer{}).Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) GetByUserId(userId int) (*model.Customer, error) {
	db := infrastructure.GetDB()

	var customer *model.Customer
	if err := db.Model(&model.Customer{}).Where("user_id = ?", userId).First(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) GetAllCall(userId int) ([]model.Call, error) {
	db := infrastructure.GetDB()

	customer, err := r.GetByUserId(userId)
	if err != nil {
		return nil, err
	}
	var calls []model.Call
	if err := db.Preload("Segments").Model(&model.Call{}).Where("customer_id = ?", customer.Id).Find(&calls).Error; err != nil {
		return nil, err
	}

	return calls, nil
}

func (r *customerRepository) FilterCallInTime(payload *model.CallTimeFilterPayload) ([]model.Call, error) {
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
func (r *customerRepository) CreateCustomer(customer *model.Customer) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Customer{}).Create(customer).Error; err != nil {
		return err
	}

	return nil

}

func (r *customerRepository) UpdateCustomer(customer *model.Customer) error {
	db := infrastructure.GetDB()

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&model.Customer{}).Updates(customer).Error; err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) GetById(id int) (*model.Customer, error) {
	db := infrastructure.GetDB()
	var customer *model.Customer

	if err := db.Model(&model.Customer{}).Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) DeleteCustomer(id int) error {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Customer{}).Where("id = ?", id).Delete(&model.Customer{}).Error; err != nil {
		return err
	}

	return nil
}
func NewCustomerRepository() model.CustomerRepository {
	return &customerRepository{}
}
