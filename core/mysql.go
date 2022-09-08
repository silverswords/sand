package core

import (
	"github.com/silverswords/sand/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDatabase struct {
	*gorm.DB
}

func NewMysqlDatabase(types, dsn string) *mysqlDatabase {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	switch types {
	case "mysql":
		session := db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin").Session(&gorm.Session{})

		session.AutoMigrate(&model.User{}, &model.Order{})

		migrator := session.Migrator()
		if !migrator.HasTable(&model.User{}) {
			migrator.CreateTable(&model.User{})
		}

		if !migrator.HasTable(&model.Order{}) {
			migrator.CreateTable(&model.Order{})
		}

		return &mysqlDatabase{
			DB: session,
		}
	default:
		panic("Unsupported Database Type")
	}
}
