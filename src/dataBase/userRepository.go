package dataBase

import (
	"context"
	"time"

	models "modules/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := MongoConnection.Database("Twittor").Collection("Users")

	user.Password, _ = EncrypPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}

func UpdateUser(user models.User, userId string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := MongoConnection.Database("Twittor").Collection("Users")

	newUser := models.User{
		Name:      user.Name,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		DateBirth: user.DateBirth,
		Banner:    user.Banner,
		Biografy:  user.Biografy,
		Location:  user.Location,
		WebSite:   user.WebSite,
	}
	Obj := bson.M{"$set": newUser}

	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := collection.UpdateOne(ctx, filter, Obj)
	if err != nil {
		return false, err
	}

	return true, nil
}
