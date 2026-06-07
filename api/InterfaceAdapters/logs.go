package InterfaceAdapters

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Registros struct {
	PilhaErro    string
	MensagemErro string
}

func EscreverLogsMongoDb(mensagem_erro string, pilha_erro string) (bool, string) {

	banco_logs, erro_conexao_banco := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://lmp:lorenzo05@kravmagaapplogs.eg0vvd9.mongodb.net/"))

	if erro_conexao_banco != nil {
		return false, "Error to connect mongodb"
	}

	banco := banco_logs.Database("kravmagalogs")

	registros := Registros{mensagem_erro, pilha_erro}
	banco.Collection("logs").InsertOne(context.Background(), registros)

	return true, "The content was write on logs"
}
