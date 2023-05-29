package models

type LoginUser struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
