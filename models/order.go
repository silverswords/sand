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
			order_id		VARCHAR(128) UNIQUE NOT NULL,
			store_id		VARCHAR(128) UNIQUE NOT NULL,
			pro_id			VARCHAR(128) UNIQUE NOT NULL,
			open_id			VARCHAR(128) UNIQUE NOT NULL,
			count			INT NOT NULL DEFAULT 1,
			total_price 	DOUBLE NOT NULL,
			status			INT NOT NULL DEFAULT 0,
			created_time 	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (order_id)
		)  ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, OrderTableName),
		fmt.Sprintf(`INSERT INTO %s (order_id, store_id, pro_id, open_id, count, total_price, status) VALUES (?, ?, ?, ?, ?, ?, ?)`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, pro_id, count, total_price, status FROM %s WHERE open_id = ?`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, pro_id, count, total_price, status FROM %s WHERE store_id = ?`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, store_id, pro_id, open_id, count, total_price, status FROM %s WHERE order_id = ?`, OrderTableName),
		fmt.Sprintf(`UPDATE %s SET status = ? WHERE order_id = ?`, OrderTableName),
	}
)

type Order struct {
	OrderID    string  `json:"order_id,omitempty"`
	ProID      string  `json:"pro_id,omitempty"`
	OpenID     string  `json:"open_id,omitempty"`
	StoreID    string  `json:"store_id,omitempty"`
	Count      uint64  `json:"count,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
	Status     uint8   `json:"status,omitempty"`
	CreateTime string  `json:"create_time"`
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
func InsertOrder(db *sql.DB, order Order) error {
	_, err := db.Exec(orderSQLString[mysqlInsertOrder], order.OrderID, order.StoreID,
		order.ProID, order.OpenID, order.Count, order.TotalPrice, order.Status)
	if err != nil {
		return err
	}

	return nil
}

// Get brife order info by user's openID
func GetOrderBrifeInfoByOpenID(db *sql.DB, openID string) ([]*Order, error) {
	rows, err := db.Query(orderSQLString[mysqlOrderBrifeInfoByOpenID], openID)
	if err != nil {
		return nil, err
	}

	var result []*Order
	for rows.Next() {
		var (
			order_id    string
			pro_id      string
			count       uint64
			total_price float64
			status      uint8
		)
		if err := rows.Scan(&order_id, &pro_id, &count, &total_price, &status); err != nil {
			return nil, err
		}

		result = append(result, &Order{
			OrderID:    order_id,
			ProID:      pro_id,
			Count:      count,
			TotalPrice: total_price,
			Status:     status,
		})
	}

	return result, nil
}

// Get brife order info by virtual store ID
func GetOrderBrifeInfoByStoreID(db *sql.DB, storeID string) ([]*Order, error) {
	rows, err := db.Query(orderSQLString[mysqlOrderBrifeInfoByStoreID], storeID)
	if err != nil {
		return nil, err
	}

	var result []*Order
	for rows.Next() {
		var (
			order_id    string
			pro_id      string
			count       uint64
			total_price float64
			status      uint8
		)
		if err := rows.Scan(&order_id, &pro_id, &count, &total_price, &status); err != nil {
			return nil, err
		}

		result = append(result, &Order{
			OrderID:    order_id,
			ProID:      pro_id,
			Count:      count,
			TotalPrice: total_price,
			Status:     status,
		})
	}

	return result, nil
}

// Get order detial info by orderID
func GetOrderDetialByOrderID(db *sql.DB, orderID string) (*Order, error) {
	var (
		order_id    string
		pro_id      string
		open_id     string
		store_id    string
		count       uint64
		total_price float64
		status      uint8
	)

	err := db.QueryRow(orderSQLString[mysqlOrderDetialInfoByOrderID], orderID).Scan(&order_id,
		&pro_id, &open_id, &store_id, &count, &total_price, &status)
	if err != nil {
		return nil, err
	}

	result := &Order{
		OrderID:    orderID,
		ProID:      pro_id,
		OpenID:     open_id,
		StoreID:    store_id,
		Count:      count,
		TotalPrice: total_price,
		Status:     status,
	}

	return result, nil
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
