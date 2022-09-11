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

type Order struct {
	OrderID    string  `gorm:"type:varchar(64);unique;not null;primaryKey"`
	UserID     string  `gorm:"type:varchar(64);not null;"`
	ProductID  string  `gorm:"type:varchar(64);not null;"`
	StoreID    string  `gorm:"type:varchar(64);not null;"`
	Quantity   uint32  `gorm:"not null"`
	TotalPrice float64 `gorm:"type:float;precision:8;scale:2;not null"`
	Status     uint8   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Product struct {
	ProductID string      `gorm:"type:varchar(64);unique;not null;primaryKey;"`
	StoreID   string      `gorm:"type:varchar(64);not null;default:'0';"`
	Price     float64     `gorm:"type:float;precision:8;scale:2;not null;default:9999.99;"`
	MainTitle string      `gorm:"type:varchar(256);"`
	Subtitle  string      `gorm:"type:varchar(256);"`
	Images    interface{} `gorm:"type:json"`
	Stock     uint32      `gorm:"not null"`
	Status    uint8       `gorm:"not null;default:0;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
