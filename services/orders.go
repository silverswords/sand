package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type orders struct {
	interfaces.DatabaseAccessor
}

func CreateOrdersService(accessor interfaces.DatabaseAccessor) Orders {
	return &orders{
		DatabaseAccessor: accessor,
	}
}

func (s *orders) Create(o *model.Order) error {
	return s.GetDefaultGormDB().Model(model.Order{}).Create(o).Error
}

func (s *orders) Modify(o *model.Order) error {
	return s.GetDefaultGormDB().Model(model.Order{}).Where("id = ?", o.ID).Updates(o).Error
}

func (s *orders) QueryByUserIDAndStatus(userID uint64, status uint8) ([]*model.Order, error) {
	var orders []*model.Order
	err := s.GetDefaultGormDB().Model(model.Order{}).
		Where("user_id = ? And status = ?", userID, status).Find(&orders).Error

	return orders, err
}
