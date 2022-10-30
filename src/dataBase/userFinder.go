package dataBase

import (
	"context"
	"time"

	models "modules/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func FindUserByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	collection := MongoConnection.Database("Twittor").Collection("Users")

	filter := bson.M{
		"email": email,
	}

	var userDb models.User

	err := collection.FindOne(ctx, filter).Decode(&userDb)
	Id := userDb.ID.Hex()
	if err != nil {
		return userDb, false, Id
	}

	return userDb, true, Id

}

func Login(email string, password string) (models.User, bool) {
	user, exist, _ := FindUserByEmail(email)

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
