package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type products struct {
	interfaces.DatabaseAccessor
}

func CreateProductsService(accessor interfaces.DatabaseAccessor) Products {
	return &products{
		DatabaseAccessor: accessor,
	}
}

func (s *products) Create(p *model.Product) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Create(p).Error
}
