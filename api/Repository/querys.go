package Repository

import (
	"fmt"
	"reflect"
	"strings"
)

func createQueryInsert(entity any, size int) (string, []any) {
	var placeholders []string
	var args []any
	table_name := strings.Split(reflect.TypeOf(entity).String(), ".")[1]

	for i := 0; i < size; i++ {
		field := reflect.ValueOf(entity).Field(i)

		switch field.Kind() {
		case reflect.Int:
			args = append(args, int(field.Int()))
		case reflect.Struct:
			args = append(args, int(field.Field(0).Int()))
		default:
			args = append(args, field.String())
		}

		placeholders = append(placeholders, "?")
	}

	query := fmt.Sprintf("INSERT INTO %s VALUES (%s)", table_name, strings.Join(placeholders, ","))
	return query, args
}

/*
 O primeiro campo do Fields será também a coluna do WHERE, passe os campos e lembre-se
 que o primeira coluna sempre será a WHERE ou seja

 USERNAME, PASSWORD WHERE USERNAME =

 2,3
*/

func createQuerySelectWhere(entity any, fields []int, content_where string) (string, []any) {
	raw_table_name := strings.Split(reflect.TypeOf(entity).String(), ".")
	table_name := raw_table_name[1]
	var fields_string []string
	var field_where string

	for i := 0; i < len(fields); i++ {
		value := fields[i]
		fields_string = append(fields_string, reflect.TypeOf(entity).Field(value).Name)
		if i == 0 {
			field_where = reflect.TypeOf(entity).Field(value).Name
		}
	}

	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s = ?",
		strings.Join(fields_string, ","),
		table_name,
		field_where,
	)

	return query, []any{content_where}
}

//SELECT USERNAME, PASSWORD WHERE USERNAME = ""
