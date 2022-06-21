package model

import "time"

type Customer struct {
	Id        int       `json:"id" gorm:"unique;autoIncrement:true"`
	Phone     string    `json:"phone" gorm:"primaryKey;column:phone"`
	UserId    int       `json:"userId" gorm:"userId"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`

	User  *User  `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE" swaggerignore:"true"`
	Calls []Call `json:"calls" gorm:"foreignKey:Phone;constraint:OnDelete:CASCADE, OnUpdate:CASCADE" swaggerignore:"true"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByUserId(userId int) (*Customer, error)
	GetAllCall(phone string) ([]Call, error)
	FilterCallInTime(payload *CallTimeFilterPayload) ([]Call, error)
	CreateCustomer(customer *Customer) error
	UpdateCustomer(customer *CustomerPhoneUpdate) error
	GetById(id int) (*Customer, error)
	DeleteCustomer(id int) error
}

type CallTimeFilterPayload struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Phone     string    `json:"phone"`
}

type CustomerPhoneUpdate struct {
	Id    int    `json:"id"`
	Phone string `json:"phone"`
}
