package model

type Employee struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"name"`
	Calls []Call `json:"calls" gorm:"foreignKey:EmployeeId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
}
