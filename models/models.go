package models

type Product struct {
	ProID      uint32      `json:"pro_id,omitempty"`
	StoreID    uint32      `json:"store_id,omitempty"` // to get every store's sales
	Price      float64     `json:"price,omitempty"`
	MainTitle  string      `json:"main_title,omitempty"`
	Subtitle   string      `json:"subtitle,omitempty"`
	Images     interface{} `json:"images,omitempty"`
	Stock      uint32      `json:"stock,omitempty"`
	Status     uint8       `json:"status,omitempty"`
	CreateTime string      `json:"create_time,omitempty"`
}

type Order struct {
	OrderID    uint32  `json:"order_id,omitempty"`
	ProID      uint32  `json:"pro_id,omitempty"`
	OpenID     uint32  `json:"open_id,omitempty"`
	StoreID    uint32  `json:"store_id,omitempty"`
	Count      uint32  `json:"count,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
	Status     uint8   `json:"status,omitempty"`
	CreateTime string  `json:"create_time"`
}
