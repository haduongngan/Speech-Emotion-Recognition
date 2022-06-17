package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id                 int    `json:"id" gorm:"primaryKey"`
	Username           string `json:"username" gorm:"unique;column:username"`
	Password           string `json:"password" gorm:"column:password"`
	Role               string `json:"role" gorm:"column:role"`
	CompanyName        string `json:"companyName" gorm:"column:companyName"`
	jwt.StandardClaims `gorm:"-" swaggerignore:"true"`
}

type UserResponse struct {
	Id           int        `json:"id"`
	Username     string     `json:"username"`
	LocationName string     `json:"locationName"`
	Role         string     `json:"role"`
	Permission   bool       `json:"permission"`
	StartTime    *time.Time `json:"startTime"`
	EndTime      *time.Time `json:"endTime"`
	Progress     int        `json:"progress"`
}

type UserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SetPermissionPayload struct {
	ReceiverUsername string `json:"receiverUsername"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
}
type CreateResponse struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	LocationName string `json:"locationName"`
	Role         string `json:"role"`
	Permission   bool   `json:"permission"`
	Message      string `json:"message"`
	Success      bool   `json:"success"`
}

type IdProgressPayload struct {
	Username string `json:"username"`
	Progress int    `json:"progress"`
}

type SexChartData struct {
	Male   int `json:"male"`
	Female int `json:"female"`
}
type AgeChartData struct {
	Kid    int
	Worker int
	Elder  int
}
type UserRepository interface {
	GetAll() ([]User, error)
	// CreateUser(user *User) (*User, error)
	// GetById(id int) (*User, error)
	// GetByUsername(username string) (*User, error)
	// DeleteUser(id int) (*User, error)
	// LoginTokenRequest(*User) (bool, error)
	// SetPermission(permission bool, receiverUsername string, startTime *time.Time, endTime *time.Time) (*User, error)
	// GetChildUser(username string) ([]User, error)
	// GetCensusProgress(username string) (interface{}, error)
	// ChangePermissionWhenLogin(username string) error
	// FalsePermissionToChild(username string) error
	// SetProgress(username string, progress int) (*User, error)
	// GetSexChart(username string) (*SexChartData, error)
	// GetChildUsernameAllLevel(username string) ([]string, error)
	// GetAgeChart(username string) (*AgeChartData, error)
	// // GetSexChartForA1() (*SexChartData, error)
}
