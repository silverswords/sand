package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type category struct {
	interfaces.DatabaseAccessor
}

func CreateCategoryService(accessor interfaces.DatabaseAccessor) Category {
	return &category{
		DatabaseAccessor: accessor,
	}
}

func (s *category) Create(c *model.Category) error {
	return s.GetDefaultGormDB().Model(model.Category{}).Create(c).Error
}

func (s *category) ModifyCategoryStatus(id uint64, status uint8) error {
	return s.GetDefaultGormDB().Model(model.Category{}).Where("id = ?", id).Update("status", status).Error
}

func (s *category) ModifyCategoryName(id uint64, name string) error {
	return s.GetDefaultGormDB().Model(model.Category{}).Where("id = ?", id).Update("name", name).Error
}

func (s *category) ListAllParentDirectory() ([]*Category, error) {
	var categories []*Category
	result := s.GetDefaultGormDB().Where("parent_id = ?", 0).Find(&categories)
	err := result.Error
	return categories, err
}

func (s *category) ListChildrenByParentID(parentID uint64) ([]*Category, error) {
	var categories []*Category
	result := s.GetDefaultGormDB().Where("parent_id = ?", parentID).Find(&categories)
	err := result.Error
	return categories, err
}
