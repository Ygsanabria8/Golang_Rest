package repository

import (
	"context"
	"time"

	dataBase "modules/dataBase"
	models "modules/src/models"
	utils "modules/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTweet(tweet *models.Tweet) (*models.Tweet, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Tweet)

	tweetDoc := bson.M{
		"userId":    tweet.UserId,
		"message":   tweet.Message,
		"createdAt": time.Now(),
	}

	result, err := collection.InsertOne(ctx, tweetDoc)
	if err != nil {
		return nil, false, err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	tweet.ID = objId

	return tweet, true, nil
}
