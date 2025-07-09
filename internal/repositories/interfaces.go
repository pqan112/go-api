package repositories

type UserRepo interface {
	FindAll()
	FindByUUID()
	Create()
	Update()
	Delete()
}
