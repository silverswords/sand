package mysql

import (
	"database/sql"
	"fmt"

	"github.com/silverswords/sand/models"
)

const OrderTableName = "order"

const (
	mysqlCreateOrderTable = iota
	mysqlInsertOrder
	mysqlOrderInfoByOpenID
	mysqlOrderInfoByStoreID
)

var (
	orderSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			order_id		BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			store_id		BIGINT UNSIGNED NOT NULL DEFAULT 0,
			pro_id			BIGINT UNSIGNED NOT NULL,
			open_id			BIGINT UNSIGNED NOT NULL,
			count			INT NOT NULL DEFAULT 1,
			total_price 	DOUBLE NOT NULL,
			status			INT NOT NULL DEFAULT 0,
			created_time 	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (order_id)
		)  ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, OrderTableName),
		fmt.Sprintf(`INSERT INTO %s (order_id, store_id, pro_id, open_id, count, total_price, status) VALUES (?, ?, ?, ?, ?, ?, ?)`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, store_id, pro_id, open_id, count, total_price, status FROM %s WHERE open_id = ?`, OrderTableName),
		fmt.Sprintf(`SELECT order_id, store_id, pro_id, open_id, count, total_price, status FROM %s WHERE store_id = ?`, OrderTableName),
	}
)

// Create order table
func CreateOrderTable(db *sql.DB) error {
	_, err := db.Exec(orderSQLString[mysqlCreateOrderTable])
	if err != nil {
		return err
	}

	return nil
}

// Insert an order
func InsertOrder(order models.Order) error {
	_, err := db.Exec(orderSQLString[mysqlInsertOrder], order.OrderID, order.StoreID,
		order.ProID, order.OpenID, order.Count, order.TotalPrice, order.Status)
	if err != nil {
		return err
	}

	return nil
}
