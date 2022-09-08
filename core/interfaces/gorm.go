package interfaces

import (
	"gorm.io/gorm"
)

type DatabaseAccessor interface {
	GetDefaultGormDB() *gorm.DB
}
