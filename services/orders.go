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
