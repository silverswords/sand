package model

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username = "root"
	password = "my-123456"
	host     = "mqying.xyz"
	port     = 3306
	Dbname   = "vs"

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, Dbname)
)

func TestCreateVSTable(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(VirtualStore{})
}

func TestCreate(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}

	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("%2d", i)

		vs := VirtualStore{
			Name:     name,
			Status:   0,
			Owner:    "my",
			Mobile:   "1593300000",
			Describe: "sell apples",
		}

		CreateVS(db, vs)
	}
}

func TestVSDelete(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}

	t.Log(DeleteVS(db, 1))
}
