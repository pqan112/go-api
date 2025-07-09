package services

type UserService interface {
	GetUsers()
	GetUserByUUID()
	CreateUser()
	UpdateUser()
}
