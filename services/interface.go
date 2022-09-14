package services

import (
	"github.com/silverswords/sand/model"
)

type Service interface {
	Users() Users
	Orders() Orders
	Products() Products
	Category() Category
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
	ListAllProducts() ([]*model.Product, error)
	QueryByProductId(id uint) (*model.Product, error)
	QueryByStoreId(storeID uint) ([]*model.Product, error)
	ModifyProduct(id uint, property string, v interface{}) error
	DeleteByProductID(id uint) error
	DeleteByStoreID(storeID uint) error
}

type Category interface {
	Create(*model.Category) error
	ChangeCategoryStatus(id uint, status int8) error
	ChangeCategoryName(id uint, name string) error
	ListChildrenByParentID(parentID uint) ([]*Category, error)
}
