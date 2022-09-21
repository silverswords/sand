package services

import (
	"github.com/silverswords/sand/model"
)

type Service interface {
	Users() Users
	Products() Products
	Category() Category
	Orders() Orders
	ShoppingCarts() ShoppingCarts
	VirtualStore() VirtualStore
	WeChat() WeChat
}

type Users interface {
	Create(*model.User) error
	QueryByOpenID(string) (*model.User, error)
	Update(*model.User) error
}

type Products interface {
	Create(*model.Product) error
	QueryDetialByProductID(id uint64) (*model.Product, error)
	QueryStockByProductID(id uint64) (uint32, error)
	ListByStoreId(storeID uint64) ([]*model.Product, error)
	ListByCategoryID(categoryID uint64) ([]*model.Product, error)
	ListAllProducts(categoryID uint64) ([]*model.Product, error)
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

type Orders interface {
	Create(*model.Order, []*model.OrderDetail) error
	Modify(*model.Order) error
	QueryByUserIDAndStatus(uint64, uint8) ([]*orderInfo, error)
	QueryDetailsByOrderID(orderID uint64) ([]*model.OrderDetail, error)
}

type ShoppingCarts interface {
	Create(*model.CartItem) error
	Query(uint64) ([]*itemInfo, error)
	Delete(itemIDs []uint64) error
	ModifyQuantity(uint64, uint32) error
}

type VirtualStore interface {
	Create(*model.VirtualStore) error
}

type Dynamic interface {
	Create(*model.Dynamic) error
}

type WeChat interface {
	Login(string) (*LoginResponse, error)
}
