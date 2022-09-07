package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UnionID string `gorm:"type:varchar(256);unique;not null"`
	OpenID  string `gorm:"type:varchar(256);unique;not null"`
	Mobile  string `gorm:"type:varchar(64) ;unique;not null"`
}

func CreateUser(db *gorm.DB, unionID string, openID string, mobile string) error {
	return db.Model(User{}).Create(&User{UnionID: unionID, OpenID: openID, Mobile: mobile}).Error
}

func QueryByUnionId(db *gorm.DB, unionID string) (uint, error) {
	var id uint
	if err := db.Model(User{}).Select("id").Where("union_id = ?", unionID).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func QueryByMobile(db *gorm.DB, mobile string) (*User, error) {
	var user User
	if err := db.Model(User{}).Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateMobile(db *gorm.DB, unionID string, mobile string) error {
	return db.Model(User{}).Where("union_id = ?", unionID).Update("mobile", mobile).Error
}
