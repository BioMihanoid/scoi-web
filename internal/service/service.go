package service

type UserServiceInterface interface {
	CreateUser()
	UpdateUser()
}

type Service struct {
	UserServiceInterface
}
