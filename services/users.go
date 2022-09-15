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
	return s.GetDefaultGormDB().Model(model.User{}).Create(u).Error
}

func (s *users) QueryByOpenID(openID string) (*model.User, error) {
	var user *model.User
	err := s.GetDefaultGormDB().Model(model.User{}).Where("open_id = ?", openID).First(&user).Error

	return user, err
}

func (s *users) Update(u *model.User) error {
	return s.GetDefaultGormDB().Model(model.User{}).Where("id = ?", u.ID).Updates(u).Error
}
