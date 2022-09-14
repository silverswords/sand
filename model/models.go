package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UnionID string `gorm:"type:varchar(256);unique;not null"`
	OpenID  string `gorm:"type:varchar(256);unique;not null"`
	Mobile  string `gorm:"type:varchar(64) ;unique;not null"`
}

type Product struct {
	gorm.Model
	StoreID    uint    `gorm:"not null;default:0"`
	CategoryID uint    `gorm:"not null"`
	Price      float64 `gorm:"precision:8;scale:2;not null"`
	PhotoUrls  string  `gorm:"type:json"`
	MainTitle  string  `gorm:"type:varchar(256)"`
	Subtitle   string  `gorm:"type:varchar(256)"`
	Status     uint8   `gorm:"type:tinyint;not null;"`
	Stock      uint32  `gorm:"not null"`
}

type Category struct {
	gorm.Model
	ParentID uint   `gorm:"not null;default:0"`
	Name     string `gorm:"not null;unique"`
	Status   uint8  `gorm:"not null;default:0"`
}

type Order struct {
	gorm.Model
	ProductID  uint      `gorm:"type:bigint;not null"`
	TotalPrice float64   `gorm:"type:decimal;precision:8;scale:2;not null"`
	UserID     uint      `gorm:"type:bigint;not null"`
	UserName   string    `gorm:"type:varchar(64);not null"`
	UserPhone  string    `gorm:"type:varchar(64);not null"`
	UserAddr   string    `gorm:"type:varchar(64);not null"`
	Status     uint8     `gorm:"type:tinyint;not null"`
	PayTime    time.Time `gorm:"type:datetime"`
}

type OrderDetail struct {
	gorm.Model
	OrderID   uint    `gorm:"type:bigint;not null"`
	ProductID uint    `gorm:"type:bigint;not null"`
	Quantity  uint    `gorm:"type:bigint;not null"`
	Price     float64 `gorm:"type:decimal;precision:8;scale:2;not null"`
}

type ShoppingCart struct {
	gorm.Model
	UserID    uint `gorm:"type:bigint;not null"`
	ProductID uint `gorm:"type:bigint;not null"`
	Quantity  uint `gorm:"type:bigint;not null"`
}

type VirtualStore struct {
	gorm.Model
	Name   string `gorm:"type:varchar(64);not null;unique"`
	Status uint8  `gorm:"type:tinyint;not null;default:0"`
}
