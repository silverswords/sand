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

func (s *users) UpdateMobile(u *model.User) error {
	return s.GetDefaultGormDB().Model(model.User{}).
		Where("union_id = ?", u.UnionID).Update("mobile", u.Mobile).Error
}
