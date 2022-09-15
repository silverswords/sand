package model

import (
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, parentID uint64, name string, status uint8) (uint64, error) {
	category := Category{ParentID: parentID, Name: name, Status: status}
	result := db.Create(&category)
	return category.ID, result.Error
}

func ChangeCategoryStatus(db *gorm.DB, id uint, status int8) error {
	return db.Model(Category{}).Where("id = ?", id).Update("status", status).Error
}

func ChangeCategoryName(db *gorm.DB, id uint, name string) error {
	return db.Model(Category{}).Where("id = ?", id).Update("name", name).Error
}

func ListChildrenByParentID(db *gorm.DB, parentID uint) ([]*Category, error) {
	var categories []*Category
	result := db.Where("parent_id = ?", parentID).Find(&categories)
	err := result.Error
	return categories, err
}
