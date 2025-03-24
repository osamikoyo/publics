package entity

type LoginRequest struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}