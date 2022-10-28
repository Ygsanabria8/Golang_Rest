package database

import (
	"context"
	"time"

	models "modules/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindUserByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := MongoConnection.Database("Twittor").Collection("Users")

	filter := bson.M{
		"email": email,
	}

	var userDb models.User

	err := collection.FindOne(ctx, filter).Decode(&userDb)
	Id := userDb.ID.Hex()
	if err != nil {
		return userDb, false, Id
	}

	return userDb, true, Id

}
