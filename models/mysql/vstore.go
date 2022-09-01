package mysql

import (
	"database/sql"
	"fmt"

	"github.com/silverswords/sand/models"
)

const VirtualStoreTableName = "vstore"

const (
	mysqlCreateVStoreTable = iota
	mysqlInsertVStore
	mysqlGetAllVirtualStores
	mysqlVStoreIsExist
)

var (
	virtualStoreSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			store_name	VARCHAR(100) NOT NULL,
			store_id	VARCHAR(20) NOT NULL,
			PRIMARY KEY (store_id)
		)  ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, VirtualStoreTableName),
		fmt.Sprintf(`INSERT INTO %s (store_name, store_id) VALUES (?, ?)`, VirtualStoreTableName),
		fmt.Sprintf(`SELECT store_name, store_id FROM %s`, VirtualStoreTableName),
	}
)

func CreateVStoreTable(db *sql.DB) error {
	_, err := db.Query(virtualStoreSQLString[mysqlCreateOrderTable])
	if err != nil {
		return err
	}

	return nil
}

func CheckVStoreIsExist(db, store_id string) {

}

// Eight nums in total, first six nums are set by admin, two nums left are random nums
func InsertVStore(db *sql.DB, vs models.VirtualStore) error {
	// rand.Seed(time.Now().UnixNano())
	// store_id := fmt.Sprintf("%s%d%d", vs.StoreID, rand.Intn(10), rand.Intn(10))

	_, err := db.Exec(virtualStoreSQLString[mysqlInsertVStore], vs.StoreName, vs.StoreID)
	if err != nil {
		return err
	}

	return nil
}

// Get all virtual stores info, only show first six nums
func GetAllVirtualStores(db *sql.DB) ([]*models.VirtualStore, error) {
	rows, err := db.Query(virtualStoreSQLString[mysqlGetAllVirtualStores])
	if err != nil {
		return nil, err
	}

	var result []*models.VirtualStore
	for rows.Next() {
		var (
			store_name string
			store_id   string
		)
		if err := rows.Scan(&store_name, &store_id); err != nil {
			return nil, err
		}
		//store_id = store_id[0:6]

		result = append(result, &models.VirtualStore{
			StoreName: store_name,
			StoreID:   store_id,
		})
	}

	return result, nil
}
