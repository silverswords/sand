package mysql

import (
	"errors"
	"fmt"
	"time"

	"github.com/silverswords/sand/models/structs"
)

const (
	TableName = "user"
)

const (
	mysqlUserCreateTable = iota
	mysqlUserInsert
	mysqlUserModifyMobile
	mysqlUserGetInfo
)

var (
	errInvalidMysql = errors.New("affected 0 rows")

	userSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			union_id    VARCHAR(128) UNIQUE NOT NULL,
			open_id     VARCHAR(128) UNIQUE NOT NULL,
			mobile   	VARCHAR(32) UNIQUE DEFAULT NULL,
			created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			modified_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (union_id)
		) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, TableName),
		fmt.Sprintf(`INSERT INTO %s (union_id, open_id)  VALUES (?,?)`, TableName),
		fmt.Sprintf(`UPDATE %s SET mobile=?, modified_at=? WHERE union_id = ? LIMIT 1`, TableName),
		fmt.Sprintf(`SELECT mobile FROM %s WHERE union_id = ? LOCK IN SHARE MODE`, TableName),
	}
)

func CreateTable() error {
	_, err := db.Exec(userSQLString[mysqlUserCreateTable])
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(union_id, open_id string) error {
	result, err := db.Exec(userSQLString[mysqlUserInsert], union_id, open_id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidMysql
	}

	return nil
}

func ModifyMobile(union_id string, time time.Time, mobile string) error {
	result, err := db.Exec(userSQLString[mysqlUserModifyMobile], mobile, time, union_id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidMysql
	}

	return nil
}

func GetUserInfo(union_id string) (*structs.User, error) {
	var (
		mobile string
	)

	err := db.QueryRow(userSQLString[mysqlUserGetInfo], union_id).Scan(&mobile)
	if err != nil {
		return nil, err
	}

	return &structs.User{
		Mobile: mobile,
	}, nil
}
