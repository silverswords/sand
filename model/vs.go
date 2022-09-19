package model

import "gorm.io/gorm"

func CreateVS(db *gorm.DB, vs VirtualStore) error {
	return db.Create(&vs).Error
}

func DeleteVS(db *gorm.DB, id uint64) error {
	result := db.Where("id = ?", id).Delete(&VirtualStore{})
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}
