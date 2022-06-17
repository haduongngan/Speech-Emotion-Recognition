package model

type Employee struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	UserId int    `json:"userId" gorm:"userId"`
	Name   string `json:"name" gorm:"name"`

	User  *User  `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
	Calls []Call `json:"calls" gorm:"foreignKey:EmployeeId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
}

type EmployeeRepository interface {
	GetAll() ([]Employee, error)
	CreateEmployee(employee *Employee) error
	UpdateEmployee(employee *Employee) error
}
