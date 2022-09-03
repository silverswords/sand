package services

type Service interface {
	Users() Users
}

type Users interface {
	Create() error
}
