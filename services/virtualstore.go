package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type virtualstore struct {
	interfaces.DatabaseAccessor
}

func CreateVirtualStoreService(accessor interfaces.DatabaseAccessor) VirtualStore {
	return &virtualstore{
		DatabaseAccessor: accessor,
	}
}

func (s virtualstore) Create(vs *model.VirtualStore) error {
	return s.GetDefaultGormDB().Model(model.VirtualStore{}).Create(vs).Error
}
