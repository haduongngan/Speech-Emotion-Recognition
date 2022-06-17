package model

import "time"

type Customer struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Phone     string    `json:"phone" gorm:"phone"`
	UserId    int       `json:"userId" gorm:"userId"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`

	User  *User  `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
	Calls []Call `json:"calls" gorm:"foreignKey:CustomerId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByUserId(userId int) (*Customer, error)
	GetAllCall(customerId int) ([]Call, error)
	FilterCallInTime(payload *CallTimeFilterPayload) ([]Call, error)
	CreateCustomer(customer *Customer) error
	UpdateCustomer(customer *Customer) error
	GetById(id int) (*Customer, error)
	DeleteCustomer(id int) error
}

type CallTimeFilterPayload struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	UserId    int       `json:"userId"`
}
