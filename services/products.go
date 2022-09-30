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

func (s *products) ProductsCreate(p *model.Product) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Create(p).Error
}

func (s *products) ProductsListAllProducts(categoryID uint64) ([]*model.Product, error) {
	var products []*model.Product
	result := s.GetDefaultGormDB().Where("status = ? AND stock > ? AND category_id = ?", 0, 0, categoryID).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) ProductsListByCategoryID(categoryID uint64) ([]*model.Product, error) {
	var products []*model.Product
	result := s.GetDefaultGormDB().Where("status = ? AND stock > ? AND category_id = ?", 0, 0, categoryID).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) ProductsQueryDetialByProductID(id uint64) (*model.Product, error) {
	var product *model.Product
	result := s.GetDefaultGormDB().Where("id = ?", id).Find(&product)
	err := result.Error
	return product, err
}

func (s *products) ProductsQueryStockByProductID(id uint64) (uint32, error) {
	var product model.Product
	err := s.GetDefaultGormDB().Select("stock").Where("id = ?", id).Find(&product).Error
	return product.Stock, err
}

func (s *products) ProductsListByStoreId(storeID uint64) ([]*model.Product, error) {
	var products []*model.Product
	result := s.GetDefaultGormDB().Where("store_id = ?", storeID).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) ProductsDeleteByProductID(id uint64) error {
	result := s.GetDefaultGormDB().Where("id = ?", id).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *products) ProductsDeleteByStoreID(storeID uint64) error {
	result := s.GetDefaultGormDB().Where("store_id = ?", storeID).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *products) ProductsModifyProduct(product *model.Product) error {
	return s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", product.ID).Updates(product).Error
}
