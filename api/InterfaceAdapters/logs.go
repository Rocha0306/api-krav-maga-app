package InterfaceAdapters

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Logs struct {
	StackTrace   string
	ErrorMessage string
}

func WriteLogsMongoDb(error_message string, StackTrace string) (bool, string) {

	log_database, connection_error_database := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://lmp:lorenzo05@kravmagaapplogs.eg0vvd9.mongodb.net/"))

	if connection_error_database != nil {
		return false, "Error to connect mongodb"
	}

	database := log_database.Database("kravmagalogs")

	Logs := Logs{error_message, StackTrace}
	database.Collection("logs").InsertOne(context.Background(), Logs)

	return true, "The content was write on logs"
}
