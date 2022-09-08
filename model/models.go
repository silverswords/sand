package model

import "gorm.io/gorm"

type Order struct {
	OrderID    uint32
	UserID     uint64
	ProductID  string
	StoreID    string
	Quantity   uint32
	TotalPrice float64
	Status     uint8
	CreateTime string
}

type User struct {
	gorm.Model
	UnionID string `gorm:"type:varchar(256);unique;not null"`
	OpenID  string `gorm:"type:varchar(256);unique;not null"`
	Mobile  string `gorm:"type:varchar(64) ;unique;not null"`
}
