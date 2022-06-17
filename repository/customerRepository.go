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

func NewCustomerRepository() model.CustomerRepository {
	return &customerRepository{}
}
