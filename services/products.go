package services

import (
	"errors"

	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

var (
	errInvalidNoRowsAffected = errors.New("affected 0 rows")
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

func (s *products) ListAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	result := s.GetDefaultGormDB().Where("status = ? AND stock > ?", 0, 0).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) ListByCategoryID(categoryID uint64) ([]*model.Product, error) {
	var products []*model.Product
	result := s.GetDefaultGormDB().Where("status = ? AND stock > ? AND category_id = ?", 0, 0, categoryID).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) QueryDetialByProductID(id uint64) (*model.Product, error) {
	var product *model.Product
	result := s.GetDefaultGormDB().Where("id = ?", id).Find(&product)
	err := result.Error
	return product, err
}

func (s *products) ListByStoreId(storeID uint64) ([]*model.Product, error) {
	var products []*model.Product
	result := s.GetDefaultGormDB().Where("store_id = ?", storeID).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) DeleteByProductID(id uint64) error {
	result := s.GetDefaultGormDB().Where("id = ?", id).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *products) DeleteByStoreID(storeID uint64) error {
	result := s.GetDefaultGormDB().Where("store_id = ?", storeID).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *products) ModifyProduct(product *model.Product) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", product.ID).Updates(product).Error
}

func (s *products) ModifyCategoryID(id []uint64, v uint64) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id IN ?", id).Updates(model.Product{CategoryID: v}).Error
}

func (s *products) ModifyStoreID(id []uint64, v uint64) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id IN ?", id).Updates(model.Product{StoreID: v}).Error
}

func (s *products) ModifyStatus(id []uint64, v uint8) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id IN ?", id).Updates(model.Product{Status: v}).Error
}
