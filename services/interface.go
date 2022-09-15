package services

import (
	"github.com/silverswords/sand/model"
)

type Service interface {
	Users() Users
	Products() Products
	Category() Category
	Orders() Orders
	OrderDetails() OrderDetails
	ShoppingCarts() ShoppingCarts
	VirtualStore() VirtualStore
}

type Users interface {
	Create(*model.User) error
	UpdateMobile(*model.User) error
}

type Products interface {
	QueryByStoreId(storeID uint8) ([]*model.Product, error)
	QueryByProductId(id uint8) (*model.Product, error)
	ListAllProducts() ([]*model.Product, error)
	Create(*model.Product) error
	ModifyCategoryID(id uint64, v interface{}) error
	ModifyPhotoUrls(id uint64, v interface{}) error
	ModifyMainTitle(id uint64, v interface{}) error
	ModifySubtitle(id uint64, v interface{}) error
	ModifyStoreID(id uint64, v interface{}) error
	ModifyStatus(id uint64, v interface{}) error
	ModifyStock(id uint64, v interface{}) error
	ModifyPrice(id uint64, v interface{}) error
	DeleteByStoreID(storeID uint8) error
	DeleteByProductID(id uint8) error
}

type Category interface {
	Create(*model.Category) error
	ChangeCategoryStatus(id uint8, status uint8) error
	ChangeCategoryName(id uint8, name string) error
	ListChildrenByParentID(parentID uint8) ([]*Category, error)
}

type VirtualStore interface {
	Create(*model.VirtualStore) error
}

type Orders interface {
	Create(*model.Order) error
}

type OrderDetails interface {
	Create(*model.OrderDetial) error
}

type ShoppingCarts interface {
	Create(*model.ShoppingCart) error
}
