package finder

import (
	"context"
	"log"
	"time"

	dataBase "modules/dataBase"
	models "modules/src/models"
	utils "modules/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweetByUserId(userId string, page int64) ([]*models.Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Tweet)

	var tweets []*models.Tweet
	filter := bson.M{"userId": userId}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "CreatedAt", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		log.Fatal(err.Error())
		return nil, false
	}

	for cursor.Next(context.TODO()) {
		var record *models.Tweet
		err := cursor.Decode(&record)
		if err != nil {
			log.Fatal(err.Error())
			return nil, false
		}
		tweets = append(tweets, record)
	}
	return tweets, true
}
