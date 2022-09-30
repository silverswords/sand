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

func (s virtualstore) VirtualStoreCreate(vs *model.VirtualStore) error {
	return s.GetDefaultGormDB().Model(model.VirtualStore{}).Create(vs).Error
}

func (s virtualstore) Delete(vs *model.VirtualStore, id uint64) error {
	result := s.GetDefaultGormDB().Model(model.VirtualStore{}).Where("id = ?", id).Delete(&vs)
	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}
