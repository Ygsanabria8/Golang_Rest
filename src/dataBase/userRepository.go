package database

import (
	"context"
	"time"

	models "modules/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)

	defer cancel()

	db := MongoConnection.Database("Twittor")
	collection := db.Collection("usuarios")

	user.Password, _ = EncrypPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
