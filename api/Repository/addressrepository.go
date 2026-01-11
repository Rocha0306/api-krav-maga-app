package Repository

import (
	entities "api-back-end/api/Entities"
)

func InsertAddress(entity entities.Address, table_name string, size int) {
	ConnectDatabase().Exec(CreateQueryInsert(entity, table_name, size))
}
