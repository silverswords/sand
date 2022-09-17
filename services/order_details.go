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

func (s *orderDetails) Create(d []*model.OrderDetail) error {
	return s.GetDefaultGormDB().Model(model.OrderDetail{}).Create(d).Error
}

func (s *orderDetails) QueryByOrderID(orderID uint64) ([]*model.OrderDetail, error) {
	var details []*model.OrderDetail
	err := s.GetDefaultGormDB().Model(model.OrderDetail{}).Where("order_id = ?", orderID).
		Find(&details).Error

	return details, err
}
