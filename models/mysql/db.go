package mysql

import (
	"database/sql"

	"github.com/silverswords/sand/sql/mysql"
)

var (
	db *sql.DB
)

func init() {
	db = mysql.CreateBuilder("./config/sql").Build().Run()
}
