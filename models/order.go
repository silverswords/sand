package models

import (
	"database/sql"
	"fmt"
)

const OrderTableName = "orders"

const (
	mysqlCreateOrderTable = iota
	mysqlInsertOrder
	mysqlOrderBrifeInfoByOpenID
	mysqlOrderBrifeInfoByStoreID
	mysqlOrderDetialInfoByOrderID
	mysqlModifyOrderStatusByOrderID
)

var (
	orderSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			order_id	 INT UNSIGNED NOT NULL AUTO_INCREMENT,
			user_id		 BIGINT UNSIGNED NOT NULL,
			pro_id 		 VARCHAR(50) NOT NULL,
			store_id 	 VARCHAR(50) NOT NULL DEFAULT '000000',
			quantity	 INT UNSIGNED NOT NULL,
			total_price  DOUBLE UNSIGNED NOT NULL,
			status 		 TINYINT UNSIGNED DEFAULT '0',
			created_at 	 DATETIME DEFAULT NOW(),
			PRIMARY KEY (order_id)
		)  ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, OrderTableName),
		fmt.Sprintf(`INSERT INTO %s (order_id, user_id, pro_id, store_id, quantity, total_price, status) VALUES (?, ?, ?, ?, ?, ?, ?)`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, pro_id, quantity, total_price, status FROM %s WHERE open_id = ?`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, pro_id, quantity, total_price, status FROM %s WHERE store_id = ?`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, user_id, pro_id, store_id, quantity, total_price, status, create_at FROM %s WHERE order_id = ?`, OrderTableName),
		fmt.Sprintf(`UPDATE %s SET status = ? WHERE order_id = ?`, OrderTableName),
	}
)

type Order struct {
	OrderID    uint32  `json:"order_id,omitempty"`
	UserID     uint64  `json:"open_id,omitempty"`
	ProductID  string  `json:"pro_id,omitempty"`
	StoreID    string  `json:"store_id,omitempty"`
	Quantity   uint32  `json:"count,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
	Status     uint8   `json:"status,omitempty"`
	CreateTime string  `json:"create_time"`
}

type Item struct {
	OrderID    uint32  `json:"order_id,omitempty"`
	ProductID  string  `json:"pro_id,omitempty"`
	Quantity   uint32  `json:"count,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
	Status     uint8   `json:"status,omitempty"`
}

// Create order table
func CreateOrderTable(db *sql.DB) error {
	_, err := db.Exec(orderSQLString[mysqlCreateOrderTable])
	if err != nil {
		return err
	}

	return nil
}

// Insert an order, get all info from admin
func (order Order) InsertOrder(db *sql.DB) error {
	_, err := db.Exec(orderSQLString[mysqlInsertOrder], order.OrderID, order.UserID,
		order.ProductID, order.StoreID, order.Quantity, order.TotalPrice, order.Status)
	if err != nil {
		return err
	}

	return nil
}

// Get brife order info by user's openID
func ListOrderByUserID(db *sql.DB, userID uint64) ([]*Item, error) {
	rows, err := db.Query(orderSQLString[mysqlOrderBrifeInfoByOpenID], userID)
	if err != nil {
		return nil, err
	}

	var result []*Item
	for rows.Next() {
		var (
			order_id    uint32
			product_id  string
			quantity    uint32
			total_price float64
			status      uint8
		)
		if err := rows.Scan(&order_id, &product_id, &quantity, &total_price, &status); err != nil {
			return nil, err
		}

		result = append(result, &Item{
			OrderID:    order_id,
			ProductID:  product_id,
			Quantity:   quantity,
			TotalPrice: total_price,
			Status:     status,
		})
	}

	return result, nil
}

// Get brife order info by virtual store ID
func ListOrderByStoreID(db *sql.DB, storeID string) ([]*Item, error) {
	rows, err := db.Query(orderSQLString[mysqlOrderBrifeInfoByStoreID], storeID)
	if err != nil {
		return nil, err
	}

	var result []*Item
	for rows.Next() {
		var (
			order_id    uint32
			product_id  string
			quantity    uint32
			total_price float64
			status      uint8
		)
		if err := rows.Scan(&order_id, &product_id, &quantity, &total_price, &status); err != nil {
			return nil, err
		}

		result = append(result, &Item{
			OrderID:    order_id,
			ProductID:  product_id,
			Quantity:   quantity,
			TotalPrice: total_price,
			Status:     status,
		})
	}

	return result, nil
}

// Get order detial info by orderID
func GetOrderDetialByOrderID(db *sql.DB, orderID string) (*Order, error) {
	return nil, nil
}

func ModifyOrderStatus(db *sql.DB, orderID string, status uint8) error {
	result, err := db.Exec(orderSQLString[mysqlModifyOrderStatusByOrderID], status, orderID)
	if err != nil {
		return nil
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidMysql
	}

	return nil
}