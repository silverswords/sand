package services

type service struct {
	users    Users
	orders   Orders
	products Products
}

func (s *service) Users() Users {
	return s.users
}

func (s *service) Orders() Orders {
	return s.orders
}

func (s *service) Products() Products {
	return s.products
}

func CreateService(u Users, o Orders, p Products) Service {
	return &service{
		users:    u,
		orders:   o,
		products: p,
	}
}
