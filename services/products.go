package services

import (
	"errors"

	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
	"gorm.io/gorm"
)

var (
	errInvalidNoRowsAffected = errors.New("affected 0 rows")
	errInvalidProperty       = errors.New("invalid property")
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

func (s *products) ListAllProducts(db *gorm.DB) ([]*model.Product, error) {
	var products []*model.Product
	result := db.Find(&products)
	err := result.Error
	return products, err
}

func (s *products) QueryByProductId(db *gorm.DB, id uint) (*model.Product, error) {
	var product *model.Product
	result := db.Where("id = ?", id).Find(&product)
	err := result.Error
	return product, err
}

func (s *products) QueryByStoreId(db *gorm.DB, storeID uint) ([]*model.Product, error) {
	var products []*model.Product
	result := db.Where("store_id = ?", storeID).Find(&products)
	err := result.Error
	return products, err
}

func (s *products) ModifyProduct(db *gorm.DB, id uint, property string, v interface{}) error {
	switch property {
	case "category_id":
		return db.Model(model.Product{}).Where("id = ?", id).Update("category_id", v).Error
	case "photo_urls":
		return db.Model(model.Product{}).Where("id = ?", id).Update("photo_urls", v).Error
	case "main_title":
		return db.Model(model.Product{}).Where("id = ?", id).Update("main_title", v).Error
	case "store_id":
		return db.Model(model.Product{}).Where("id = ?", id).Update("store_id", v).Error
	case "subtitle":
		return db.Model(model.Product{}).Where("id = ?", id).Update("subtitle", v).Error
	case "status":
		return db.Model(model.Product{}).Where("id = ?", id).Update("status", v).Error
	case "stock":
		return db.Model(model.Product{}).Where("id = ?", id).Update("stock", v).Error
	case "price":
		return db.Model(model.Product{}).Where("id = ?", id).Update("price", v).Error
	}

	return errInvalidProperty
}

func (s *products) DeleteByProductID(db *gorm.DB, id uint) error {
	result := db.Where("id = ?", id).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *products) DeleteByStoreID(db *gorm.DB, storeID uint) error {
	result := db.Where("store_id = ?", storeID).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}
