package repository

import (
	"context"
	dataBase "modules/dataBase"
	"modules/src/models"
	utils "modules/src/utils"
	"time"
)

func CreateFollow(follow *models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Follow)

	_, err := collection.InsertOne(ctx, follow)
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteFollow(follow *models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Follow)

	_, err := collection.DeleteOne(ctx, follow)
	if err != nil {
		return false, err
	}

	return true, nil
}
