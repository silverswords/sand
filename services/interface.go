package services

import "github.com/silverswords/sand/model"

type Service interface {
	Users() Users
}

type Users interface {
	Create(*model.User) error
	UpdateMobile(*model.User) error
}
