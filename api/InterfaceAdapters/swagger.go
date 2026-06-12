package InterfaceAdapters

import (
	"strings"
)

// Documentacao da API (aberta). O Controller (camada Presentation) chama
// estas funcoes; aqui fica so o conteudo (HTML do Swagger UI + spec OpenAPI).

// SwaggerUIHTML devolve o HTML do Swagger UI ja apontando pro openapi.json.
func SwaggerUIHTML() string {
	return strings.ReplaceAll(swaggerUIHTML, "__OPENAPI_URL__", "/internal/docs/openapi.json")
}

// OpenApiSpec devolve o spec OpenAPI 3.0 da API.
func OpenApiSpec() string {
	return openapiSpec
}

const swaggerUIHTML = `<!DOCTYPE html>
<html lang="pt-br">
<head>
  <meta charset="UTF-8">
  <title>KravConnect API</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.ui = SwaggerUIBundle({
      url: "__OPENAPI_URL__",
      dom_id: "#swagger-ui",
      presets: [SwaggerUIBundle.presets.apis],
      layout: "BaseLayout"
    });
  </script>
</body>
</html>`

const openapiSpec = `{
  "openapi": "3.0.3",
  "info": {
    "title": "KravConnect API",
    "version": "1.0.0",
    "description": "API do app de Krav Maga. Endpoints protegidos exigem header Authorization: Bearer <token>."
  },
  "servers": [{ "url": "/" }],
  "components": {
    "securitySchemes": {
      "bearerAuth": { "type": "http", "scheme": "bearer", "bearerFormat": "JWT" }
    },
    "schemas": {
      "LoginDTO": { "type": "object", "required": ["email","senha"], "properties": { "email": {"type":"string","format":"email"}, "senha": {"type":"string"} } },
      "UsuarioDTO": { "type": "object", "required": ["email","senha","cpf","cep"], "properties": { "email": {"type":"string","format":"email"}, "senha": {"type":"string","maxLength":20}, "cpf": {"type":"string","minLength":11,"maxLength":11}, "cep": {"type":"string","minLength":8,"maxLength":8} } },
      "CodigoUsuarioDTO": { "type": "object", "required": ["codigo_auth"], "properties": { "codigo_auth": {"type":"string","maxLength":6} } },
      "EsqueciSenhaDTO": { "type": "object", "required": ["email"], "properties": { "email": {"type":"string","format":"email"} } },
      "RedefinirSenhaDTO": { "type": "object", "required": ["codigo_auth","nova_senha"], "properties": { "codigo_auth": {"type":"string","maxLength":6}, "nova_senha": {"type":"string","maxLength":20} } },
      "AcademiaDTO": { "type": "object", "required": ["cnpj","nome"], "properties": { "cnpj": {"type":"string","minLength":14,"maxLength":14}, "nome": {"type":"string","maxLength":100} } },
      "CodigoConviteDTO": { "type": "object", "required": ["convite_uuid"], "properties": { "convite_uuid": {"type":"string","format":"uuid"} } },
      "AprovarSolicitacaoDTO": { "type": "object", "required": ["id_solicitacao"], "properties": { "id_solicitacao": {"type":"string","maxLength":100} } },
      "AtualizarFaixaDTO": { "type": "object", "required": ["faixa"], "properties": { "faixa": {"type":"string","maxLength":15} } },
      "CriarInstrutorDTO": { "type": "object", "required": ["id_aluno"], "properties": { "id_aluno": {"type":"string","maxLength":100} } },
      "CriarAulaDTO": { "type": "object", "required": ["conteudo","data_aula","faixa"], "properties": { "conteudo": {"type":"string","maxLength":255}, "data_aula": {"type":"string","example":"2026-06-11 19:30"}, "faixa": {"type":"string","maxLength":15} } },
      "AtualizarAulaDTO": { "type": "object", "required": ["id_aula","conteudo","data_aula","faixa"], "properties": { "id_aula": {"type":"string"}, "conteudo": {"type":"string","maxLength":255}, "data_aula": {"type":"string","example":"2026-06-11 19:30"}, "faixa": {"type":"string","maxLength":15} } },
      "PresencaDTO": { "type": "object", "required": ["id_aula","latitude","longitude"], "properties": { "id_aula": {"type":"string","maxLength":100}, "latitude": {"type":"number","format":"double"}, "longitude": {"type":"number","format":"double"} } },
      "LocalizacaoAcademiaDTO": { "type": "object", "required": ["latitude","longitude"], "properties": { "latitude": {"type":"number","format":"double"}, "longitude": {"type":"number","format":"double"} } },
      "CriarProdutoDTO": { "type": "object", "required": ["nome","preco","tamanho","quantidade"], "properties": { "nome": {"type":"string","maxLength":100}, "preco": {"type":"number","format":"double","maximum":5000}, "tamanho": {"type":"string","maxLength":20}, "quantidade": {"type":"integer","minimum":1,"maximum":20}, "imagem_url": {"type":"string","format":"uri","maxLength":500} } },
      "AtualizarProdutoDTO": { "type": "object", "required": ["id_produto","nome","preco","tamanho","quantidade"], "properties": { "id_produto": {"type":"string","maxLength":100}, "nome": {"type":"string","maxLength":100}, "preco": {"type":"number","format":"double","maximum":5000}, "tamanho": {"type":"string","maxLength":20}, "quantidade": {"type":"integer","minimum":1,"maximum":20}, "imagem_url": {"type":"string","format":"uri","maxLength":500} } },
      "SinalizarInteresseDTO": { "type": "object", "required": ["id_produto","quantidade"], "properties": { "id_produto": {"type":"string","maxLength":100}, "quantidade": {"type":"integer","minimum":1} } },
      "PagamentoDTO": { "type": "object", "required": ["valor_centavos"], "properties": { "valor_centavos": {"type":"integer","format":"int64","minimum":100} } },
      "EnviarEmailDTO": { "type": "object", "required": ["para","conteudo"], "properties": { "para": {"type":"string","format":"email"}, "conteudo": {"type":"string"} } }
    }
  },
  "security": [{ "bearerAuth": [] }],
  "paths": {
    "/Users/Auth": { "post": { "tags": ["Users"], "summary": "Login", "security": [], "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/LoginDTO" } } } }, "responses": { "200": { "description": "Token JWT" }, "400": { "description": "Credenciais invalidas" } } } },
    "/Users/Me": { "get": { "tags": ["Users"], "summary": "Perfil do usuario logado", "responses": { "200": { "description": "Perfil" } } } },
    "/Users/Registration": { "post": { "tags": ["Users"], "summary": "Cadastro (envia codigo por email)", "security": [], "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/UsuarioDTO" } } } }, "responses": { "200": { "description": "Codigo enviado" } } } },
    "/Users/Registration/Confirm": { "post": { "tags": ["Users"], "summary": "Confirma cadastro com codigo", "security": [], "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/CodigoUsuarioDTO" } } } }, "responses": { "200": { "description": "Usuario criado" } } } },
    "/Users/Password/Forgot": { "post": { "tags": ["Users"], "summary": "Solicita codigo de redefinicao de senha", "security": [], "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/EsqueciSenhaDTO" } } } }, "responses": { "200": { "description": "OK" } } } },
    "/Users/Password/Reset": { "post": { "tags": ["Users"], "summary": "Redefine a senha com codigo", "security": [], "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/RedefinirSenhaDTO" } } } }, "responses": { "200": { "description": "Senha alterada" } } } },
    "/Gyms/Creation": { "post": { "tags": ["Gyms"], "summary": "Cria academia (vira professor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/AcademiaDTO" } } } }, "responses": { "200": { "description": "Academia criada" } } } },
    "/Gyms/List": { "get": { "tags": ["Gyms"], "summary": "Lista academias do usuario (professor e aluno)", "responses": { "200": { "description": "Lista de vinculos" } } } },
    "/Gyms/Invites/Creation": { "post": { "tags": ["Gyms"], "summary": "Gera convite (professor)", "responses": { "200": { "description": "Convite criado" } } } },
    "/Gyms/Invites/List": { "get": { "tags": ["Gyms"], "summary": "Lista convites da academia (professor)", "responses": { "200": { "description": "Lista" } } } },
    "/Gyms/Invites/{id_convite}": { "delete": { "tags": ["Gyms"], "summary": "Deleta convite (professor)", "parameters": [{ "name": "id_convite", "in": "path", "required": true, "schema": { "type": "string" } }], "responses": { "200": { "description": "Removido" } } } },
    "/Gyms/Requests/Join": { "post": { "tags": ["Gyms"], "summary": "Solicita entrada via convite", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/CodigoConviteDTO" } } } }, "responses": { "200": { "description": "Solicitacao registrada" } } } },
    "/Gyms/Requests/Join/List": { "get": { "tags": ["Gyms"], "summary": "Lista solicitacoes de entrada (professor)", "responses": { "200": { "description": "Lista" } } } },
    "/Gyms/Requests/Join/Approve": { "post": { "tags": ["Gyms"], "summary": "Aprova solicitacao (professor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/AprovarSolicitacaoDTO" } } } }, "responses": { "200": { "description": "Aprovado" } } } },
    "/Gyms/Students/List": { "get": { "tags": ["Gyms"], "summary": "Lista alunos da academia (professor)", "responses": { "200": { "description": "Lista" } } } },
    "/Gyms/Students/{id_aluno}": { "delete": { "tags": ["Gyms"], "summary": "Remove aluno (professor)", "parameters": [{ "name": "id_aluno", "in": "path", "required": true, "schema": { "type": "string" } }], "responses": { "200": { "description": "Removido" } } } },
    "/Gyms/Students/{id_aluno}/Belt": { "put": { "tags": ["Gyms"], "summary": "Atualiza faixa do aluno (professor)", "parameters": [{ "name": "id_aluno", "in": "path", "required": true, "schema": { "type": "string" } }], "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/AtualizarFaixaDTO" } } } }, "responses": { "200": { "description": "Atualizado" } } } },
    "/Gyms/Instructors/Creation": { "post": { "tags": ["Gyms"], "summary": "Promove aluno a instrutor (professor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/CriarInstrutorDTO" } } } }, "responses": { "200": { "description": "Instrutor criado" } } } },
    "/Gyms/Classes/Creation": { "post": { "tags": ["Classes"], "summary": "Cria aula (professor/instrutor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/CriarAulaDTO" } } } }, "responses": { "200": { "description": "Aula criada" } } } },
    "/Gyms/Classes/Update": { "put": { "tags": ["Classes"], "summary": "Atualiza aula (professor/instrutor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/AtualizarAulaDTO" } } } }, "responses": { "200": { "description": "Aula atualizada" } } } },
    "/Gyms/Classes/Day": { "get": { "tags": ["Classes"], "summary": "Aulas do dia", "responses": { "200": { "description": "Lista" } } } },
    "/Gyms/Classes": { "get": { "tags": ["Classes"], "summary": "Aulas da academia (professor)", "parameters": [{ "name": "data", "in": "query", "required": false, "schema": { "type": "string", "example": "2026-06-11" } }], "responses": { "200": { "description": "Lista" } } } },
    "/Student/Presence": { "post": { "tags": ["Student"], "summary": "Registra presenca (geolocalizada)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/PresencaDTO" } } } }, "responses": { "200": { "description": "Presenca registrada" } } } },
    "/Student/Presence/Count": { "get": { "tags": ["Student"], "summary": "Presencas de um aluno (professor)", "parameters": [{ "name": "id_aluno", "in": "query", "required": true, "schema": { "type": "string" } }], "responses": { "200": { "description": "Contagem" } } } },
    "/Gyms/Location": { "put": { "tags": ["Gyms"], "summary": "Define localizacao da academia (professor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/LocalizacaoAcademiaDTO" } } } }, "responses": { "200": { "description": "OK" } } } },
    "/Gyms/Geolocation": { "get": { "tags": ["Gyms"], "summary": "Localizacao da academia do usuario", "responses": { "200": { "description": "Lat/Long" } } } },
    "/Gyms/Catalog/Creation": { "post": { "tags": ["Catalog"], "summary": "Cria produto (professor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/CriarProdutoDTO" } } } }, "responses": { "200": { "description": "Produto criado" } } } },
    "/Gyms/Catalog/Update": { "put": { "tags": ["Catalog"], "summary": "Atualiza produto (professor)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/AtualizarProdutoDTO" } } } }, "responses": { "200": { "description": "Produto atualizado" } } } },
    "/Gyms/Catalog/{id_produto}": { "delete": { "tags": ["Catalog"], "summary": "Deleta produto (professor)", "parameters": [{ "name": "id_produto", "in": "path", "required": true, "schema": { "type": "string" } }], "responses": { "200": { "description": "Removido" } } } },
    "/Gyms/Catalog": { "get": { "tags": ["Catalog"], "summary": "Lista produtos de uma academia", "parameters": [{ "name": "id_academia", "in": "query", "required": true, "schema": { "type": "string" } }], "responses": { "200": { "description": "Lista" } } } },
    "/Student/Interest": { "post": { "tags": ["Student"], "summary": "Sinaliza interesse em produto (aluno)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/SinalizarInteresseDTO" } } } }, "responses": { "200": { "description": "Interesse registrado" } } } },
    "/Gyms/Products/Email": { "post": { "tags": ["Catalog"], "summary": "Envia email (Brevo)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/EnviarEmailDTO" } } } }, "responses": { "200": { "description": "Email enviado" } } } },
    "/Gyms/Pix/Onboarding": { "post": { "tags": ["Payments"], "summary": "Onboarding Stripe Connect (professor): retorna URL de cadastro Pix", "responses": { "200": { "description": "URL de onboarding" } } } },
    "/Student/Payment": { "post": { "tags": ["Payments"], "summary": "Cria pagamento Pix (aluno)", "requestBody": { "required": true, "content": { "application/json": { "schema": { "$ref": "#/components/schemas/PagamentoDTO" } } } }, "responses": { "200": { "description": "client_secret do PaymentIntent" } } } }
  }
}`
