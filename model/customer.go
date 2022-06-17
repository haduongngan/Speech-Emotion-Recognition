package model

type Customer struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Phone  string `json:"phone" gorm:"phone"`
	UserId int    `json:"userId" gorm:"userId"`

	User  *User  `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
	Calls []Call `json:"calls" gorm:"foreignKey:CustomerId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	CreateCustomer(customer *Customer) error
	UpdateCustomer(customer *Customer) error
}
