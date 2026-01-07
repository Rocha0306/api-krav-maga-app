package InterfaceAdapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// 450 requisicoes por mes
var api_keys = [9]string{
	"2afee4e70dd514eb8bfbe2083d9bafed85b5e116f1e108bce0042fb034ceadf4",
	"194e09caa373f0338ed45592c9b4215f3f2f17af0ce6da120aab0dbb59afa132",
	"8e972dfdaff01be25ceca3e962d1f061cbdfa8f8ee11b17a36637878690a474c",
	"759e94ae2d4d23bf5fe25c36a0527e43196845093011053ffd34062adcd142ef",
	"e869bad067fe337ecc584c884f9966a201d01a9d50a25bdf417de42a51ee5037",
	"260212add0c3ddc6f21e0a3c36b68f0bc909c044a35075e469ef73ebad2aeb4e",
	"432b8c32c9258326b9852a358c588c7aaf0327f78c07e30f106e3c23fe9d6ce6",
	"8fa316757c13db66b1d4a10f6d21a0211e68651209f762281dd9afedee661c57",
	"2e1e23409686cbd7d3fd9ee84bfdf986367606d52d5413385fc89ff1e48c0cf2",
}

type data struct {
	cpf       string
	name      string
	gender    string
	birthDate string
}

type error_api struct {
	message string
}

type CpfApiResponse struct {
	sucess    bool
	error_api error_api
	data      data
	CPF       string
	Name      string
	Gender    string
	BirthDate string
}

type CepApiResponse struct {
	CEP         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	UF          string
	Estado      string
	Regiao      string
	DDD         string
}

func CpfApi(cpf string) (error, CpfApiResponse) {
	var private_response_desserialized CpfApiResponse = CpfApiResponse{}
	url := fmt.Sprintf("https://api.cpfhub.io/cpf/%s", cpf)
	client := &http.Client{}
	for i := 0; i < 9; i++ {
		request, err := http.NewRequest("GET", url, nil)
		if err == nil {
			request.Header.Add("x-api-key", api_keys[i])
		}

		response, err := client.Do(request)

		json.NewDecoder(response.Body).Decode(&private_response_desserialized)

		if private_response_desserialized.sucess == false && private_response_desserialized.error_api.message != "" {
			return errors.New(private_response_desserialized.error_api.message), private_response_desserialized
		} else {
			return err, filteredcpfresponse(private_response_desserialized)

		}

	}

	WriteLogsMongoDb("URGENTE - O limite de requisicoes foi atingido no Cpfhub - 500", "")
	return errors.New("Problema interno no servidor"), private_response_desserialized
}

func CepApi(cep string) (error, CepApiResponse) {
	url := fmt.Sprintf("viacep.com.br/ws/%s/json/", cep)
	response_desserialized := CepApiResponse{}

	response, err := http.Get(url)

	if err != nil {
		WriteLogsMongoDb("Erro API via Cep, verificar!", "InterfaceAdapters/apis.go/CepApi()")
		return errors.New("Erro interno do servidor - 500"), response_desserialized
	}

	json.NewDecoder(response.Body).Decode(&response_desserialized)

	if err != nil {
		return err, CepApiResponse{}
	}

	return nil, response_desserialized

}

func filteredcpfresponse(raw_response CpfApiResponse) CpfApiResponse {
	var filtered_response CpfApiResponse = CpfApiResponse{}
	filtered_response.Name = raw_response.data.name
	filtered_response.CPF = raw_response.data.cpf
	filtered_response.Gender = raw_response.data.gender
	filtered_response.BirthDate = raw_response.data.birthDate
	return filtered_response
}
