package model

import "time"

type Staff struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"userId" gorm:"userId"`
	Name      string    `json:"name" gorm:"name"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`

	User  *User  `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE" swaggerignore:"true"`
	Calls []Call `json:"calls" gorm:"foreignKey:StaffId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE" swaggerignore:"true"`
}

type StaffRepository interface {
	GetAll() ([]Staff, error)
	GetByUserId(userId int) (*Staff, error)
	GetAllCall(userId int) ([]Call, error)
	FilterCallInTime(payload *StaffCallFilterPayload) ([]Call, error)
	CreateStaff(staff *Staff) error
	UpdateStaff(staffName *StaffNameUpdate) error
	GetById(id int) (*Staff, error)
	DeleteStaff(id int) error
}

type StaffNameUpdate struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
