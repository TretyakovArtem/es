package driver

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	goqu "gopkg.in/doug-martin/goqu.v5"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/mysql"
)

func ConnectDB() *goqu.Database {
	msDb, err := sql.Open("mysql", "root:17121991@/test_db")

	if err != nil {
		panic(err.Error())
	}

	db := goqu.New("mysql", msDb)

	return db
}
