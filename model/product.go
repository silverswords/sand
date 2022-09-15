package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	errInvalidNoRowsAffected = errors.New("affected 0 rows")
	errInvalidProperty       = errors.New("invalid property")
)

// Insert a product into the table
func CreateProduct(db *gorm.DB, storeID uint64, categoryID uint64, price float64,
	photoUrls string, mainTitle string, subtitle string, status uint8, stock uint32) (uint64, error) {
	product := Product{
		StoreID:    storeID,
		CategoryID: categoryID,
		Price:      price,
		PhotoUrls:  photoUrls,
		MainTitle:  mainTitle,
		Subtitle:   subtitle,
		Status:     status,
		Stock:      stock,
	}
	result := db.Create(&product)

	return product.ID, result.Error
}

// List all products after judgement product status, stock and it's virtual store
func ListAllProducts(db *gorm.DB) ([]*Product, error) {
	var products []*Product
	result := db.Where("status = ? AND stock > ?", 0, 0).Find(&products)
	err := result.Error
	return products, err
}

// Get product detial info by id
func QueryByProductId(db *gorm.DB, id uint) (*Product, error) {
	var product *Product
	result := db.Where("id = ?", id).Find(&product)
	err := result.Error
	return product, err
}

// Get virtual store's all products
func QueryByStoreId(db *gorm.DB, storeID uint) ([]*Product, error) {
	var products []*Product
	result := db.Where("store_id = ?", storeID).Find(&products)
	err := result.Error
	return products, err
}

// Modify product property
func ModifyProduct(db *gorm.DB, id uint, property string, v interface{}) error {
	switch property {
	case "category_id":
		return db.Model(Product{}).Where("id = ?", id).Update("category_id", v).Error
	case "photo_urls":
		return db.Model(Product{}).Where("id = ?", id).Update("photo_urls", v).Error
	case "main_title":
		return db.Model(Product{}).Where("id = ?", id).Update("main_title", v).Error
	case "store_id":
		return db.Model(Product{}).Where("id = ?", id).Update("store_id", v).Error
	case "subtitle":
		return db.Model(Product{}).Where("id = ?", id).Update("subtitle", v).Error
	case "status":
		return db.Model(Product{}).Where("id = ?", id).Update("status", v).Error
	case "stock":
		return db.Model(Product{}).Where("id = ?", id).Update("stock", v).Error
	case "price":
		return db.Model(Product{}).Where("id = ?", id).Update("price", v).Error
	}

	return errInvalidProperty
}

// Delete product by product ID
func DeleteByProductID(db *gorm.DB, id uint) error {
	result := db.Where("id = ?", id).Delete(&Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

// Delete all products in virtual store
func DeleteByStoreID(db *gorm.DB, storeID uint) error {
	result := db.Where("store_id = ?", storeID).Delete(&Product{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}
