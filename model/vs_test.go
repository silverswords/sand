package model

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateVS(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(VirtualStore{})
}
