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

func (s *products) QueryByProductId(id uint8) (*model.Product, error) {
	var product *model.Product
	result := s.GetDefaultGormDB().Where("id = ?", id).Find(&product)
	err := result.Error
	return product, err
}

func (s *products) QueryByStoreId(storeID uint8) ([]*model.Product, error) {
	var products []*model.Product
	result := s.GetDefaultGormDB().Where("store_id = ?", storeID).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) DeleteByProductID(id uint8) error {
	result := s.GetDefaultGormDB().Where("id = ?", id).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *products) DeleteByStoreID(storeID uint8) error {
	result := s.GetDefaultGormDB().Where("store_id = ?", storeID).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *products) ModifyCategoryID(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("category_id", v).Error
}

func (s *products) ModifyPhotoUrls(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("photo_urls", v).Error
}

func (s *products) ModifyMainTitle(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("main_title", v).Error
}

func (s *products) ModifyStoreID(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("store_id", v).Error
}

func (s *products) ModifySubtitle(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("subtitle", v).Error
}

func (s *products) ModifyStatus(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("status", v).Error
}

func (s *products) ModifyStock(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("stock", v).Error
}

func (s *products) ModifyPrice(id uint64, v interface{}) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", id).Update("price", v).Error
}
