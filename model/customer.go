package model

type Customer struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Phone string `json:"phone" gorm:"phone"`
	Calls []Call `json:"calls" gorm:"foreignKey:CustomerId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
}
