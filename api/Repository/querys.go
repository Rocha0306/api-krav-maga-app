package Repository

import (
	"fmt"
	"reflect"
	"strings"
)

func createQueryInsert(entity any, table_name string, size int) string {
	var values []string

	for i := 0; i < size; i++ {

		value := reflect.ValueOf(entity).Field(i).String()
		kind_value := reflect.ValueOf(entity).Field(i).Kind().String()

		switch kind_value {
		case "int":
			values = append(values, fmt.Sprint(int(reflect.ValueOf(entity).Field(i).Int())))
		case "struct":
			values = append(values, fmt.Sprint(int(reflect.ValueOf(entity).Field(i).Field(0).Int())))
		default:
			changed_value := fmt.Sprintf("'%s'", value)
			values = append(values, changed_value)
		}

	}

	query := fmt.Sprint(
		"INSERT INTO ",
		table_name,
		" VALUES (",
		strings.Join(values, ","),
		")",
	)

	return query

}

/*
 O primeiro campo do Fields será também a coluna do WHERE, passe os campos e lembre-se
 que o primeira coluna sempre será a WHERE ou seja

 USERNAME, PASSWORD WHERE USERNAME =

 2,3
*/

func createQuerySelectWhere(entity any, fields []int, content_where string) string {
	raw_table_name := strings.Split(reflect.TypeOf(entity).String(), ".")
	table_name := raw_table_name[1]
	var fields_string []string
	var field_where string

	//2,3,8

	for i := 0; i < len(fields); i++ {
		value := fields[i]
		if i == 0 {
			fields_string = append(fields_string, reflect.TypeOf(entity).Field(value).Name)
			field_where = reflect.TypeOf(entity).Field(value).Name
		} else {
			fields_string = append(fields_string, reflect.TypeOf(entity).Field(value).Name)
		}
	}

	query := fmt.Sprint(
		"SELECT ",
		strings.Join(fields_string, ","),
		" FROM ",
		table_name,
		" WHERE ",
		field_where,
		"=",
		fmt.Sprintf("'%s'", content_where),
	)

	return query
}

//SELECT USERNAME, PASSWORD WHERE USERNAME = ""
