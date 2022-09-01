package models

type Product struct {
	ProID      uint64      `json:"pro_id,omitempty"`
	StoreID    uint64      `json:"store_id,omitempty"` // to get every store's sales
	Price      float64     `json:"price,omitempty"`
	MainTitle  string      `json:"main_title,omitempty"`
	Subtitle   string      `json:"subtitle,omitempty"`
	Images     interface{} `json:"images,omitempty"`
	Stock      uint64      `json:"stock,omitempty"`
	Status     uint8       `json:"status,omitempty"`
	CreateTime string      `json:"create_time,omitempty"`
}

type Order struct {
	OrderID    uint64  `json:"order_id,omitempty"`
	ProID      uint64  `json:"pro_id,omitempty"`
	OpenID     uint64  `json:"open_id,omitempty"`
	StoreID    uint64  `json:"store_id,omitempty"`
	Count      uint64  `json:"count,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
	Status     uint8   `json:"status,omitempty"`
	CreateTime string  `json:"create_time"`
}

type VirtualStore struct {
	StoreName string `json:"store_name"`
	StoreID   string `json:"store_id"`
}
