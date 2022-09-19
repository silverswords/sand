package core

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
	Dbname   = "demo"

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, Dbname)
)

type PriceTable struct {
	ID    uint64
	Info  string
	Price float64 `gorm:"precision:8;scale:2;not null"`
}

func TestCreateTables(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database, error: " + err.Error())
	}

	db.AutoMigrate(PriceTable{})
}

// func TestFloatNums(t *testing.T) {
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to the database, error: " + err.Error())
// 	}

// 	type Price struct {
// 		Price float32 `gorm:"type:float;precision:8;scale:2"`
// 	}

// 	db.AutoMigrate(Price{})

// 	rows, err := db.Find(&Price{}).Rows()
// 	if err != nil {
// 		t.Log(err)
// 	}

// 	prices := make([]float32, 0)
// 	for rows.Next() {
// 		var price float32
// 		if err := rows.Scan(&price); err != nil {
// 			t.Log(err)
// 		}

// 		prices = append(prices, price)
// 	}

// 	if err := rows.Err(); err != nil {
// 		t.Log(err)
// 	}

// 	t.Log(prices)
// }
