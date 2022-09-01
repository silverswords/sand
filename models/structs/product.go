package structs

type Product struct {
	ProID      string      `json:"pro_id,omitempty"`
	StoreID    string      `json:"store_id,omitempty"` // to get every store's sales
	Price      float64     `json:"price,omitempty"`
	MainTitle  string      `json:"main_title,omitempty"`
	Subtitle   string      `json:"subtitle,omitempty"`
	Images     interface{} `json:"images,omitempty"`
	Stock      uint64      `json:"stock,omitempty"`
	Status     uint8       `json:"status,omitempty"`
	CreateTime string      `json:"create_time,omitempty"`
}

// Homepage will show brife product info
type BrifeProduct struct {
	ProID    string      `json:"pro_id,omitempty"`
	Price    float64     `json:"price,omitempty"`
	Subtitle string      `json:"subtitle,omitempty"`
	Images   interface{} `json:"images,omitempty"`
	Status   uint8       `json:"status,omitempty"`
}
