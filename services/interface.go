package services

import (
	"github.com/silverswords/sand/model"
)

type Service interface {
	Users
	Products
	Category
	Orders
	Carts
	VirtualStore
	WeChat
	Sign
}

type Users interface {
	UsersCreate(*model.User) error
	UsersQueryByID(uint64) (*model.User, error)
	UsersQueryByOpenID(string) (*model.User, error)
	UsersUpdate(*model.User) error
}

type Products interface {
	ProductsCreate(*model.Product) error
	ProductsQueryDetialByProductID(uint64) (*model.Product, error)
	ProductsQueryStockByProductID(uint64) (uint32, error)
	ProductsListByStoreId(uint64) ([]*model.Product, error)
	ProductsListByCategoryID(uint64) ([]*model.Product, error)
	ProductsListAllProducts(uint64) ([]*model.Product, error)
	ProductsModifyProduct(*model.Product) error
	ProductsDeleteByStoreID(uint64) error
	ProductsDeleteByProductID(uint64) error
}

type Category interface {
	CategoryCreate(*model.Category) error
	CategoryModifyCategoryStatus(uint64, uint8) error
	CategoryModifyCategoryName(uint64, string) error
	CategoryListAllParentDirectory() ([]*Category, error)
	CategoryListChildrenByParentID(uint64) ([]*Category, error)
}

type Orders interface {
	OrdersCreate(*model.Order, []*model.OrderDetail, *model.UserAddress, bool) error
	OrdersQueryByUserIDAndStatus(uint64, uint8) ([]*orderInfo, error)
	OrdersQueryDetailsByOrderID(uint64, uint64) (*orderDetail, error)
	OrdersModifyStatus(uint64, uint64, uint8) error
	OrdersModifyAddress(uint64, uint64, *model.UserAddress) error
	OrdersDelete(uint64, uint64) error
}

type Carts interface {
	CartsCreate(*model.CartItem) error
	CartsQuery(uint64) ([]*itemInfo, error)
	CartsDelete(uint64, []uint64) error
	CartsModifyQuantity(uint64, uint64, uint32) error
}

type VirtualStore interface {
	VirtualStoreCreate(*model.VirtualStore) error
}

type Dynamic interface {
	DynamicCreate(*model.Dynamic) error
}

type WeChat interface {
	Login(string) (*LoginResponse, error)
	GetAccessToken() (string, error)
	GetPhoneNumber(string) (*PhoneResp, error)
	GetPrepayInfo(string, string, int, string) (string, string, error)
}

type Sign interface {
	GetSignedInfo(string, string) (*PayInfo, error)
}
