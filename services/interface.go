package services

import (
	"github.com/silverswords/sand/model"
	"gorm.io/gorm"
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
	ListAllProducts(db *gorm.DB) ([]*model.Product, error)
	QueryByProductId(db *gorm.DB, id uint) (*model.Product, error)
	QueryByStoreId(db *gorm.DB, storeID uint) ([]*model.Product, error)
	ModifyProduct(db *gorm.DB, id uint, property string, v interface{}) error
	DeleteByProductID(db *gorm.DB, id uint) error
	DeleteByStoreID(db *gorm.DB, storeID uint) error
}

type Category interface {
	Create(*model.Category) error
	ChangeCategoryStatus(id uint, status int8) error
	ChangeCategoryName(id uint, name string) error
	ListChildrenByParentID(parentID uint) ([]*Category, error)
}
