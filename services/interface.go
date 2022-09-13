package services

import "github.com/silverswords/sand/model"

type Service interface {
	Users() Users
	Orders() Orders
	Products() Products
}

type Users interface {
	Create(*model.User) error
	UpdateMobile(*model.User) error
}

type Orders interface {
	Create(*model.Order) error
}

type Products interface {
	Create(*model.Product) error
}
