package dataBase

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection object connection to the data base
var MongoConnection = ConnectMongoDb()
var clientOptions = options.Client().ApplyURI("mongodb+srv://go_rest:Technic1136@bdlearning.okxfo5v.mongodb.net/?retryWrites=true&w=majority")

// ConnectMongoDb make connection with database
func ConnectMongoDb() *mongo.Client {
	clientOptions.SetMaxPoolSize(50)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Cannot connect with mongo db %v", err.Error())
		return client
	}

	log.Println("Conected to mongo data base")
	return client
}

// CheckConnection verify connection to data base
func CheckConnection() bool {
	err := MongoConnection.Ping(context.TODO(), nil)
	return err != nil
}
