package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type users struct {
	interfaces.DatabaseAccessor
}

func CreateUsersService(accessor interfaces.DatabaseAccessor) Users {
	return &users{
		DatabaseAccessor: accessor,
	}
}

func (s *users) Create(u *model.User) error {
	return nil
}
