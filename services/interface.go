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
	Create(*model.Product) error
	QueryDetialByProductID(id uint64) (*model.Product, error)
	ListByStoreId(storeID uint64) ([]*model.Product, error)
	ListByCategoryID(categoryID uint64) ([]*model.Product, error)
	ListAllProducts() ([]*model.Product, error)
	ModifyCategoryID(id []uint64, v uint64) error
	ModifyStoreID(id []uint64, v uint64) error
	ModifyStatus(id []uint64, v uint8) error
	ModifyProduct(product *model.Product) error
	DeleteByStoreID(storeID uint64) error
	DeleteByProductID(id uint64) error
}

type Category interface {
	Create(*model.Category) error
	ModifyCategoryStatus(id uint64, status uint8) error
	ModifyCategoryName(id uint64, name string) error
	ListAllParentDirectory() ([]*Category, error)
	ListChildrenByParentID(parentID uint64) ([]*Category, error)
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
