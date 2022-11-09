package finder

import (
	"context"
	dataBase "modules/dataBase"
	models "modules/src/models"
	utils "modules/src/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func FindFollow(follow *models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	filter := bson.M{
		"userId":       follow.UserId,
		"userFollowId": follow.UserFollowId,
	}

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Follow)

	var followDB *models.Follow

	err := collection.FindOne(ctx, filter).Decode(&followDB)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetTweetsFollowers(userId string, page int64) ([]models.TweetFollower, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Follow)

	skip := (page - 1) * 20

	filter := make([]bson.M, 0)

	filter = append(filter, bson.M{"$match": bson.M{"userId": userId}})
	filter = append(filter, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userFollowId",
			"foreignField": "userId",
			"as":           "tweet",
		},
	})
	filter = append(filter, bson.M{"$unwind": "$tweet"})
	filter = append(filter, bson.M{"$dort": bson.M{"CreatedAt:": -1}})
	filter = append(filter, bson.M{"$skip": skip})
	filter = append(filter, bson.M{"$limit": 20})

	cursor, err := collection.Aggregate(ctx, filter)
	var result []models.TweetFollower
	err = cursor.All(ctx, result)
	if err != nil {
		return nil, false
	}
	return result, true
}
