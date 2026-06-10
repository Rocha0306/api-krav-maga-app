package InterfaceAdapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

type CpfApiResponse struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
}

type CepApiResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	DDD         string `json:"ddd"`
}

type CnpjApiResponse struct {
}

const stripe_secret_key = "sk_test_sua_chave_stripe"

type StripePaymentIntentResponse struct {
	ID           string `json:"id"`
	ClientSecret string `json:"client_secret"`
	Status       string `json:"status"`
	Amount       int64  `json:"amount"`
	Currency     string `json:"currency"`
}

func StripeCriarPagamento(valor_centavos int64, descricao string) (StripePaymentIntentResponse, error) {
	var resposta StripePaymentIntentResponse

	dados := url.Values{}
	dados.Set("amount", strconv.FormatInt(valor_centavos, 10))
	dados.Set("currency", "brl")
	dados.Set("description", descricao)
	dados.Add("payment_method_types[]", "card")
	dados.Add("payment_method_types[]", "pix")

	request, err := http.NewRequest("POST", "https://api.stripe.com/v1/payment_intents", strings.NewReader(dados.Encode()))
	if err != nil {
		return resposta, errors.New("erro ao preparar pagamento")
	}
	request.SetBasicAuth(stripe_secret_key, "")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return resposta, errors.New("erro ao conectar com a Stripe")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return resposta, errors.New("erro ao criar pagamento na Stripe")
	}

	json.NewDecoder(response.Body).Decode(&resposta)
	return resposta, nil
}

func CpfApi(cpf string) (error, CpfApiResponse) {
	var private_response_desserialized CpfApiResponse
	url := fmt.Sprintf("https://api.cpfhub.io/cpf/%s", cpf)
	client := &http.Client{}
	for i := 0; i < 9; i++ {
		request, err := http.NewRequest("GET", url, nil)
		if err == nil {
			request.Header.Add("x-api-key", api_keys[i])
			request.Header.Add("content-type", "application/json")
			request.Header.Add("Accept-Encoding", "identity")
		}

		raw_response, err := client.Do(request)

		if raw_response.StatusCode == 400 || raw_response.StatusCode == 422 {
			return errors.New("CPF invalido"), private_response_desserialized
		}

		if raw_response.StatusCode == 403 {
			continue
		}

		json.NewDecoder(raw_response.Body).Decode(&private_response_desserialized)

		return nil, private_response_desserialized

	}

	EscreverLogsMongoDb("URGENTE - O limite de requisicoes foi atingido no Cpfhub - 500", "")
	return errors.New("Problema interno no servidor"), private_response_desserialized
}

func CepApi(cep string) (error, CepApiResponse) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	response_desserialized := CepApiResponse{}

	request, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		EscreverLogsMongoDb("Erro API via Cep, verificar!", "InterfaceAdapters/apis.go/CepApi()")
		return errors.New("Erro interno do servidor - 500"), response_desserialized
	}

	json.NewDecoder(response.Body).Decode(&response_desserialized)

	if response_desserialized.CEP == "" {
		return errors.New("CEP Invalido"), response_desserialized
	}

	return nil, response_desserialized

}
