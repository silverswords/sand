package model

import "time"

type Model struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	Model
	ParentID uint64 `gorm:"not null;default:0"`
	Name     string `gorm:"not null;unique"`
	Status   uint8  `gorm:"not null;default:0"`
}

type Order struct {
	Model
	UserID        uint64    `gorm:"not null"`
	UserAddressID uint64    `gorm:"not null"`
	TotalPrice    float64   `gorm:"precision:8;scale:2;not null"`
	Status        uint8     `gorm:"not null;default:0"`
	PayTime       time.Time `gorm:"default:null"`
}

type OrderDetail struct {
	Model
	OrderID   uint64  `gorm:"not null"`
	ProductID uint64  `gorm:"not null"`
	Quantity  uint32  `gorm:"not null"`
	Price     float64 `gorm:"precision:8;scale:2;not null"`
}

type Product struct {
	Model
	StoreID    uint64  `gorm:"not null;default:0"`
	CategoryID uint64  `gorm:"not null"`
	Price      float64 `gorm:"precision:8;scale:2;not null"`
	PhotoUrls  string  `gorm:"type:json"`
	MainTitle  string  `gorm:"type:varchar(256)"`
	Subtitle   string  `gorm:"type:varchar(256)"`
	Status     uint8   `gorm:"not null"`
	Stock      uint32  `gorm:"not null"`
}

type ShoppingCart struct {
	Model
	UserID    uint64 `gorm:"not null"`
	ProductID uint64 `gorm:"not null"`
	Quantity  uint32 `gorm:"not null"`
}

type UserAddress struct {
	Model
	UserID       uint64 `gorm:"not null"`
	UserName     string `gorm:"type:varchar(64);not null"`
	UserPhone    string `gorm:"type:varchar(64);not null"`
	ProvinceName string `gorm:"type:varchar(64);not null"`
	CityName     string `gorm:"type:varchar(64);not null"`
	CountName    string `gorm:"type:varchar(64);not null"`
	DetialInfo   string `gorm:"type:varchar(256);not null"`
}

type User struct {
	Model
	UnionID string `gorm:"type:varchar(256);unique;not null"`
	OpenID  string `gorm:"type:varchar(256);unique;not null"`
	Mobile  string `gorm:"type:varchar(64);unique;not null"`
}

type VirtualStore struct {
	Model
	Name     string `gorm:"not null;unique"`
	Status   uint8  `gorm:"not null;default:0"`
	Owner    string `gorm:"type:varchar(64);not null"`
	Mobile   string `gorm:"type:varchar(64);not null"`
	Describe string `gorm:"not null"`
}

type Dynamic struct {
	Model
	Text      string `gorm:"not null"`
	PhotoUrls string `gorm:"type:json"`
}
