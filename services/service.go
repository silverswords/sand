package services

type service struct {
	Users
	Products
	Category
	Orders
	Carts
	VirtualStore
	WeChat
	Sign
}

func CreateService(u Users, p Products, c Category, o Orders, s Carts, vs VirtualStore, w WeChat, sg Sign) Service {
	return &service{
		Users:        u,
		Orders:       o,
		Products:     p,
		Category:     c,
		Carts:        s,
		VirtualStore: vs,
		WeChat:       w,
		Sign:         sg,
	}
}
