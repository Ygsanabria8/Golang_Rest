package finder

import (
	"context"
	"log"
	"time"

	dataBase "modules/dataBase"
	models "modules/src/models"
	utils "modules/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetFollowedUsers(userId string, page int64, search string, typeUser string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Users)

	var resulst []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	filter := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err.Error())
		return nil, false
	}

	var found, include bool

	for cursor.Next(context.TODO()) {
		var record *models.User
		err := cursor.Decode(&record)
		if err != nil {
			log.Fatal(err.Error())
			return nil, false
		}

		var follow = &models.Follow{
			UserId:       userId,
			UserFollowId: record.ID.Hex(),
		}

		include = false

		found, _ = FindFollow(follow)
		if typeUser == "new" && !found {
			include = true
		}
		if typeUser == "follow" && found {
			include = true
		}

		if follow.UserFollowId == userId {
			include = false
		}

		if include {
			record.CleanUnnecessaryData()
			resulst = append(resulst, record)
		}

		resulst = append(resulst, record)
	}

	err = cursor.Err()
	if err != nil {
		log.Fatal(err.Error())
		return nil, false
	}

	cursor.Close(ctx)
	return resulst, true
}
