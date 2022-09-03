package models

import (
	"database/sql"
	"errors"
	"fmt"
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
		fmt.Sprintf(`INSERT INTO %s (union_id, open_id, mobile)  VALUES (?,?,?)`, TableName),
		fmt.Sprintf(`UPDATE %s SET mobile=?, modified_at=? WHERE union_id = ? LIMIT 1`, TableName),
		fmt.Sprintf(`SELECT mobile FROM %s WHERE union_id = ? LOCK IN SHARE MODE`, TableName),
	}
)

type User struct {
	UnionID string
	OpenID  string
	Mobile  string
}

func CreateUserTable(db *sql.DB) error {
	_, err := db.Exec(userSQLString[mysqlUserCreateTable])
	if err != nil {
		return err
	}

	return nil
}

func InsertUser(db *sql.DB, user User) error {
	result, err := db.Exec(userSQLString[mysqlUserInsert], user.UnionID, user.OpenID, user.Mobile)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidMysql
	}

	return nil
}
