package structs

type Order struct {
	OrderID    string  `json:"order_id,omitempty"`
	ProID      string  `json:"pro_id,omitempty"`
	OpenID     string  `json:"open_id,omitempty"`
	StoreID    string  `json:"store_id,omitempty"`
	Count      uint64  `json:"count,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
	Status     uint8   `json:"status,omitempty"`
	CreateTime string  `json:"create_time"`
}
