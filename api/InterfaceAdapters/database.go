package InterfaceAdapters

import (
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

const connection_string string = ""
const User string = ""
const Password string = ""
const database_name string = ""

func openconnection() (driver.Connector, error) {
	configs_mysql := mysql.NewConfig()
	configs_mysql.Addr = connection_string
	configs_mysql.User = User
	configs_mysql.Passwd = Password
	configs_mysql.DBName = database_name
	driver_connected, err := mysql.NewConnector(configs_mysql)

	if err != nil {
		WriteLogsMongoDb("Critico - Nao foi possivel se conectar ao banco",
			"/InterfaceAdapters/database.go - 500")
		return nil, err
	}

	return driver_connected, nil
}

func InsertStudent(entity any) (bool, error) {
	return tur

}
