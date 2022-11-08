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
