package dataBase

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	utils "modules/src/utils"
)

// MongoConnection object connection to the data base
var MongoConnection *mongo.Client

// ConnectMongoDb make connection with database
func ConnectMongoDb() {
	var clientOptions = options.Client().ApplyURI(utils.Config.Mongo.ConnectionString)
	clientOptions.SetMaxPoolSize(50)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Cannot connect with mongo db %v", err.Error())
		return
	}

	log.Println("Conected to mongo data base")
	MongoConnection = client
}

// CheckConnection verify connection to data base
func CheckConnection() bool {
	err := MongoConnection.Ping(context.TODO(), nil)
	return err != nil
}
