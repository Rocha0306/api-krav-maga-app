package Repository

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

const Address string = "localhost"
const User string = "root"
const Password string = "lorenzo05*"
const database_name string = "KRAVMAGAAPP"

func connectDatabase() *sql.DB {
	configs_mysql := mysql.NewConfig()
	configs_mysql.Addr = Address
	configs_mysql.User = User
	configs_mysql.Passwd = Password
	configs_mysql.DBName = database_name
	driver_connected, _ := mysql.NewConnector(configs_mysql)
	driver_connected.Connect(context.Background())
	connection_string := configs_mysql.FormatDSN()
	db, _ := sql.Open("mysql", connection_string)
	return db
}

func Insert(entity any, table_name string, size int) {
	connectDatabase().Exec(createQueryInsert(entity, table_name, size))
}

func SelectWhere(entity any, fields []int, content_where string) any {
	result, err := connectDatabase().Query(createQuerySelectWhere(entity, fields, content_where))
	if err == nil {

		for result.Next() {
			result.Scan(entity)
			return entity
		}

	} else {
		return entity
	}
	return entity
}
