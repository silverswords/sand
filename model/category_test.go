package model

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	username = "root"
// 	password = "my-123456"
// 	host     = "mqying.xyz"
// 	port     = 3306
// 	Dbname   = "vs"

// 	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		username, password, host, port, Dbname)
// )

func TestCreateCategory(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Category{})

	id, err := CreateCategory(db, 1, "水产", 0)
	if err != nil {
		t.Log(err)
	}
	t.Log(id)
}

func TestChangeCategoryStatus(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}

	err = ChangeCategoryStatus(db, 4, 1)
	if err != nil {
		t.Log(err)
	}
}

func TestChangeCategoryName(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}

	err = ChangeCategoryName(db, 8, "水果蔬菜")
	t.Log(err != nil)
}

func TestListChildrenByParentID(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}

	categories, err := ListChildrenByParentID(db, 1)
	if err != nil {
		t.Log(err)
	}

	for _, v := range categories {
		t.Log(v)
	}
}
