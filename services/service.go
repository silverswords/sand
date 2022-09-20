package services

type service struct {
	users         Users
	products      Products
	category      Category
	orders        Orders
	shoppingCarts ShoppingCarts
	virtualstore  VirtualStore
}

func (s *service) Users() Users {
	return s.users
}

func (s *service) Products() Products {
	return s.products
}

func (s *service) Category() Category {
	return s.category
}

func (s *service) Orders() Orders {
	return s.orders
}

func (s *service) ShoppingCarts() ShoppingCarts {
	return s.shoppingCarts
}

func (s *service) VirtualStore() VirtualStore {
	return s.virtualstore
}

func CreateService(u Users, p Products, c Category, o Orders, s ShoppingCarts, vs VirtualStore) Service {
	return &service{
		users:         u,
		orders:        o,
		products:      p,
		category:      c,
		shoppingCarts: s,
		virtualstore:  vs,
	}
}
