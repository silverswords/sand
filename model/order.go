package model

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
