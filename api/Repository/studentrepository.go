package Repository

import (
	entities "api-back-end/api/Entities"
)

func InsertStudent(entity entities.Student, table_name string, size int) {
	ConnectDatabase().Exec(CreateQueryInsert(entity, table_name, size))
}

func SelectOneStudent(entity any, fields []int, content_where string) entities.Student {
	student := entities.Student{}
	result, err := ConnectDatabase().Query(CreateQuerySelectWhere(entity, fields, content_where))
	if err == nil {

		for result.Next() {
			result.Scan(&student.Username, &student.Password, &student.Role)
			return student
		}

	} else {
		return student
	}
	return student
}
