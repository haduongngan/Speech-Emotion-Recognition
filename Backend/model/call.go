package model

import "time"

type Call struct {
	Id              int       `json:"id" gorm:"primaryKey"`
	Phone           string    `json:"phone" gorm:"phone"`
	StaffId         int       `json:"staffId" gorm:"staffId"`
	StartTime       time.Time `json:"startTime" gorm:"startTime"`
	Duration        string    `json:"duration" gorm:"duration"`
	StaffEmotion    string    `json:"staffEmotion" gorm:"staffEmotion"`
	CustomerEmotion string    `json:"emotion" gorm:"emotion"`
	Segments        []Segment `json:"segments" gorm:"foreignKey:CallId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE" swaggerignore:"true"`
}

type CallRepository interface {
	GetAll() ([]Call, error)
	CreateCall(call *Call) error
	CreateMultiCall(calls []Call) error
	UpdateCall(call *Call) error
	GetById(id int) (*Call, error)
	DeleteCall(id int) error
}

type HistoryPayload struct {
	Phone string `json:"phone"`
}

type StaffCallFilterPayload struct {
	UserId    int       `json:"userId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
