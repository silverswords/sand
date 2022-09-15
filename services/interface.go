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
	QueryByOpenID(openID string) (*model.User, error)
	Update(*model.User) error
}

type Products interface {
	QueryByStoreId(storeID uint8) ([]*model.Product, error)
	QueryByProductId(id uint8) (*model.Product, error)
	ListAllProducts() ([]*model.Product, error)
	Create(*model.Product) error
	ModifyCategoryID(id []uint64, v uint64) error
	ModifyStoreID(id []uint64, v uint64) error
	ModifyStatus(id []uint64, v uint8) error
	ModifyProduct(product *model.Product) error
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
	Create(*model.OrderDetail) error
}

type ShoppingCarts interface {
	Create(*model.ShoppingCart) error
}
