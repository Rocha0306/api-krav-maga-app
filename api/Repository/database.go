package Repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const Address string = "localhost"
const User string = "root"
const Password string = "lorenzo05*"
const database_name string = "KRAVMAGAAPP"

func connectDatabase() *gorm.DB {
	dsn := "root:lorenzo05*@tcp(localhost)/KRAVMAGAAPP?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
	return db

}

func Insert[T any](entity *T) {
	connectDatabase().Create(entity)
}

func SelectWhere[T any](field string, value string) *T {
	var result T
	connectDatabase().Where(field+" = ?", value).First(&result)
	return &result
}

func SelectWhereList[T any](field string, value string) []T {
	var results []T
	connectDatabase().Where(field+" = ?", value).Find(&results)
	return results
}

func SelectJoin[T any](join string, field string, value string) []T {
	var results []T
	connectDatabase().Joins(join).Where(field+" = ?", value).Find(&results)
	return results
}

func Exists[T any](field string, value string) bool {
	var model T
	var count int64
	connectDatabase().Model(&model).Where(field+" = ?", value).Count(&count)
	return count > 0
}
