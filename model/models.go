package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UnionID string `gorm:"type:varchar(256);unique;not null"`
	OpenID  string `gorm:"type:varchar(256);unique;not null"`
	Mobile  string `gorm:"type:varchar(64) ;unique;not null"`
}

type Order struct {
	gorm.Model
	OrderID    string  `gorm:"type:varchar(64);not null;primaryKey"`
	UserID     string  `gorm:"type:varchar(64);not null"`
	ProductID  string  `gorm:"type:varchar(64);not null"`
	StoreID    string  `gorm:"type:varchar(64);not null"`
	Quantity   uint32  `gorm:"not null"`
	TotalPrice float64 `gorm:"precision:8;scale:2;not null"`
	Status     uint8   `gorm:"not null"`
}

type Product struct {
	gorm.Model
	StoreID    uint    `gorm:"not null;default:0"`
	CategoryID uint    `gorm:"not null"`
	Price      float64 `gorm:"precision:8;scale:2;not null"`
	PhotoUrls  string  `gorm:"type:json"`
	MainTitle  string  `gorm:"type:varchar(256)"`
	Subtitle   string  `gorm:"type:varchar(256)"`
	Status     uint8   `gorm:"not null;default:0"`
	Stock      uint32  `gorm:"not null"`
}

type Category struct {
	gorm.Model
	ParentID uint   `gorm:"not null;default:0"`
	Name     string `gorm:"not null;unique"`
	Status   int8   `gorm:"not null;default:0"`
}
