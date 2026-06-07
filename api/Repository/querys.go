package Repository

import (
	"fmt"
	"reflect"
	"strings"
)

func criarQueryInsert(entidade any, tamanho int) (string, []any) {
	var placeholders []string
	var args []any
	nome_tabela := strings.Split(reflect.TypeOf(entidade).String(), ".")[1]

	for i := 0; i < tamanho; i++ {
		campo := reflect.ValueOf(entidade).Field(i)

		switch campo.Kind() {
		case reflect.Int:
			args = append(args, int(campo.Int()))
		case reflect.Struct:
			args = append(args, int(campo.Field(0).Int()))
		default:
			args = append(args, campo.String())
		}

		placeholders = append(placeholders, "?")
	}

	query := fmt.Sprintf("INSERT INTO %s VALUES (%s)", nome_tabela, strings.Join(placeholders, ","))
	return query, args
}

/*
 O primeiro campo do Campos sera tambem a coluna do WHERE, passe os campos e lembre-se
 que a primeira coluna sempre sera a WHERE ou seja

 USERNAME, PASSWORD WHERE USERNAME =

 2,3
*/

func criarQuerySelectWhere(entidade any, campos []int, conteudo_where string) (string, []any) {
	nome_tabela_bruto := strings.Split(reflect.TypeOf(entidade).String(), ".")
	nome_tabela := nome_tabela_bruto[1]
	var campos_string []string
	var campo_where string

	for i := 0; i < len(campos); i++ {
		valor := campos[i]
		campos_string = append(campos_string, reflect.TypeOf(entidade).Field(valor).Name)
		if i == 0 {
			campo_where = reflect.TypeOf(entidade).Field(valor).Name
		}
	}

	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s = ?",
		strings.Join(campos_string, ","),
		nome_tabela,
		campo_where,
	)

	return query, []any{conteudo_where}
}
