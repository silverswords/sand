package mysql

import (
	"fmt"

	"github.com/silverswords/sand/models/structs"
)

const VirtualStoreTableName = "vstore"

const (
	mysqlCreateVStoreTable = iota
	mysqlInsertVStore
	mysqlGetAllVirtualStores
)

var (
	virtualStoreSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			store_name	VARCHAR(100) NOT NULL UNIQUE,
			store_id	VARCHAR(20) NOT NULL,
			PRIMARY KEY (store_id)
		)  ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, VirtualStoreTableName),
		fmt.Sprintf(`INSERT INTO %s (store_name, store_id) VALUES (?, ?)`, VirtualStoreTableName),
		fmt.Sprintf(`SELECT store_name, store_id FROM %s`, VirtualStoreTableName),
	}
)

func CreateVStoreTable() error {
	_, err := db.Query(virtualStoreSQLString[mysqlCreateOrderTable])
	if err != nil {
		return err
	}

	return nil
}

// Create a new virtual store, get all info from admin
func InsertVStore(vs structs.VirtualStore) error {
	_, err := db.Exec(virtualStoreSQLString[mysqlInsertVStore], vs.StoreName, vs.StoreID)
	if err != nil {
		return err
	}

	return nil
}

// Get all virtual stores info, show them to the admin
func GetAllVirtualStores() ([]*structs.VirtualStore, error) {
	rows, err := db.Query(virtualStoreSQLString[mysqlGetAllVirtualStores])
	if err != nil {
		return nil, err
	}

	var result []*structs.VirtualStore
	for rows.Next() {
		var (
			store_name string
			store_id   string
		)
		if err := rows.Scan(&store_name, &store_id); err != nil {
			return nil, err
		}

		result = append(result, &structs.VirtualStore{
			StoreName: store_name,
			StoreID:   store_id,
		})
	}

	return result, nil
}
