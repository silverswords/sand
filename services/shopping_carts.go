package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type shoppingCarts struct {
	interfaces.DatabaseAccessor
}

func CreateShoppingCartsService(accessor interfaces.DatabaseAccessor) ShoppingCarts {
	return &shoppingCarts{
		DatabaseAccessor: accessor,
	}
}

func (s *shoppingCarts) Create(sc *model.ShoppingCart) error {
	return s.GetDefaultGormDB().Model(model.Order{}).Create(sc).Error
}
