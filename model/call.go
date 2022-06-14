package model

import "time"

type Call struct {
	Id              int        `json:"id" gorm:"primaryKey"`
	CustomerId      int        `json:"customerId" gorm:"customerId"`
	EmployeeId      int        `json:"employeeId" gorm:"employeeId"`
	StartTime       *time.Time `json:"startTime" gorm:"startTime"`
	Duration        string     `json:"duration" gorm:"duration"`
	EmployeeEmotion string     `json:"employeeEmotion" gorm:"employeeEmotion"`
	CustomerEmotion string     `json:"customerEmotion" gorm:"customerEmotion"`
	OverallEmotion  string     `json:"overallEmotion" gorm:"overallEmotion"`
	Segments        []Segment  `json:"segments" gorm:"foreignKey:CallId;constraint:OnDelete:CASCADE, OnUpdate:CASCADE"`
}
