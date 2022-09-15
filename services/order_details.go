package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type orderDetails struct {
	interfaces.DatabaseAccessor
}

func CreateOrderDetailsService(accessor interfaces.DatabaseAccessor) OrderDetails {
	return &orderDetails{
		DatabaseAccessor: accessor,
	}
}

func (s *orderDetails) Create(d *model.OrderDetail) error {
	return s.GetDefaultGormDB().Model(model.OrderDetail{}).Create(d).Error
}
