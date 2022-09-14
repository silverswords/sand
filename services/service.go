package services

type service struct {
	users    Users
	orders   Orders
	products Products
	category Category
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

func (s *service) Category() Category {
	return s.category
}

func CreateService(u Users, o Orders, p Products, c Category) Service {
	return &service{
		users:    u,
		orders:   o,
		products: p,
		category: c,
	}
}
