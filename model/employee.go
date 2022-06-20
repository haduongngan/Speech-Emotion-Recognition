package model

import "time"

type Employee struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"userId" gorm:"userId"`
	Name      string    `json:"name" gorm:"name"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`

	User  *User  `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
	Calls []Call `json:"calls" gorm:"foreignKey:EmployeeId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
}

type EmployeeRepository interface {
	GetAll() ([]Employee, error)
	GetByUserId(userId int) (*Employee, error)
	GetAllCall(userId int) ([]Call, error)
	FilterCallInTime(payload *CallTimeFilterPayload) ([]Call, error)
	CreateEmployee(employee *Employee) error
	UpdateEmployee(employee *Employee) error
	GetById(id int) (*Employee, error)
	DeleteEmployee(id int) error
}
