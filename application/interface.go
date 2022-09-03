package application

import (
	"github.com/silverswords/sand/services"
	"gorm.io/gorm"
)

type DataSource interface {
	GetDefaultGormDB() *gorm.DB
}

type ServiceProvider interface {
	Services() services.Service
}
