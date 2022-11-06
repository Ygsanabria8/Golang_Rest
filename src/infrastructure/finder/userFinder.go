package finder

import (
	"context"
	"time"

	dataBase "modules/dataBase"
	models "modules/src/models"
	utils "modules/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUserByEmail(email string) (*models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Users)

	filter := bson.M{
		"email": email,
	}

	var userDb *models.User

	err := collection.FindOne(ctx, filter).Decode(&userDb)
	if err != nil {
		return userDb, false, string("")
	}

	Id := userDb.ID.Hex()
	return userDb, true, Id

}

func FindUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Users)

	var userDb *models.User
	objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objId,
	}

	err := collection.FindOne(ctx, filter).Decode(&userDb)
	userDb.Password = ""
	if err != nil {
		return nil, err
	}

	return userDb, nil

}
