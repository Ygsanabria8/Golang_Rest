package repository

import (
	"context"
	"time"

	dataBase "modules/dataBase"
	finder "modules/src/infrastructure/finder"
	models "modules/src/models"
	utils "modules/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Users)

	user.Password, _ = utils.EncrypPassword(user.Password)

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

	collection := dataBase.MongoConnection.Database(utils.Config.Mongo.Database).Collection(utils.Config.Mongo.Users)

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

func Login(email string, password string) (models.User, bool) {
	user, exist, _ := finder.FindUserByEmail(email)

	if !exist {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDb := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDb, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
