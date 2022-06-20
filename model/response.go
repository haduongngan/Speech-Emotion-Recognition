package model

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	UserId       int    `json:"userId"`
	Role         string `json:"role"`
	Username     string `json:"username"`
	CompanyName  string `json:"companyName"`
	Message      string `json:"message"`
	Code         string `json:"code"`
	Success      bool   `json:"success"`
}
