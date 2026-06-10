package Presentation

import "api-back-end/api/InterfaceAdapters"

func DominioEmailValido(email string) bool {
	return InterfaceAdapters.RegexDominioEmail(email)
}

func CnpjValido(cnpj string) bool {
	digitos := make([]int, 0, 14)
	for _, r := range cnpj {
		if r >= '0' && r <= '9' {
			digitos = append(digitos, int(r-'0'))
		}
	}
	if len(digitos) != 14 {
		return false
	}

	todosIguais := true
	for _, d := range digitos {
		if d != digitos[0] {
			todosIguais = false
			break
		}
	}
	if todosIguais {
		return false
	}

	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	digitoVerificador := func(pesos []int) int {
		soma := 0
		for i, p := range pesos {
			soma += digitos[i] * p
		}
		resto := soma % 11
		if resto < 2 {
			return 0
		}
		return 11 - resto
	}

	return digitoVerificador(pesos1) == digitos[12] && digitoVerificador(pesos2) == digitos[13]
}
