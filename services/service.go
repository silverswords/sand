package services

type service struct {
	users Users
}

func (s *service) Users() Users {
	return s.users
}

func CreateService(u Users) Service {
	return &service{
		users: u,
	}
}
