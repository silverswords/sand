package structs

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
