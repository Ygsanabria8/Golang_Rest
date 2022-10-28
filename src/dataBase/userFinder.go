package database

import (
	"context"
	"time"

	models "modules/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindUserByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)

	defer cancel()

	db := MongoConnection.Database("Twittor")
	collection := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var userDb models.User

	err := collection.FindOne(ctx, condition).Decode(&userDb)
	Id := userDb.ID.Hex()
	if err != nil {
		return userDb, false, Id
	}

	return userDb, false, Id

}
