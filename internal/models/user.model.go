package models

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Status   string `json:"status" binding:"required,userStatus"`
}
