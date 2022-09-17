package model

import (
	"errors"

	"gorm.io/gorm"
)

var (
	errInvalidNoRowsAffected = errors.New("affected 0 rows")
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

func ListByCategoryID(db *gorm.DB, categoryID uint64) ([]*Product, error) {
	var products []*Product
	result := db.Where("status = ? AND stock > ? AND category_id = ?", 0, 0, categoryID).Find(&products)
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

func ModifyCategoryID(db *gorm.DB, id []uint64, v uint64) error {
	return db.Model(Product{}).Where("id IN ?", id).Updates(Product{CategoryID: v}).Error
}

func ModifyPhotoUrls(db *gorm.DB, id uint64, v interface{}) error {
	return db.Model(Product{}).Where("id = ?", id).Update("photo_urls", v).Error
}

func ModifyMainTitle(db *gorm.DB, id uint64, v interface{}) error {
	return db.Model(Product{}).Where("id = ?", id).Update("main_title", v).Error
}

func ModifyStoreID(db *gorm.DB, id []uint64, v uint64) error {
	return db.Model(Product{}).Where("id IN ?", id).Updates(Product{StoreID: v}).Error
}

func ModifySubtitle(db *gorm.DB, id uint64, v interface{}) error {
	return db.Model(Product{}).Where("id = ?", id).Update("subtitle", v).Error
}

func ModifyStatus(db *gorm.DB, id []uint64, v uint8) error {
	return db.Model(Product{}).Where("id IN ?", id).Updates(Product{Status: v}).Error
}

func ModifyStock(db *gorm.DB, id []uint64, v uint32) error {
	return db.Model(Product{}).Where("id IN ?", id).Updates(Product{Stock: v}).Error
}

func ModifyPrice(db *gorm.DB, id uint64, v interface{}) error {
	return db.Model(Product{}).Where("id = ?", id).Update("price", v).Error
}

func ModifyProduct(db *gorm.DB, product *Product) error {
	return db.Model(Product{}).Where("id = ?", product.ID).Updates(product).Error
}
