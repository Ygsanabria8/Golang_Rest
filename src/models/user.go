package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User type of object for a user
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	LastName  string             `bson:"lastName,omitempty" json:"lastName"`
	DateBirth time.Time          `bson:"dateBirth,omitempty" json:"dateBirth"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	Avatar    string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Banner    string             `bson:"banner,omitempty" json:"banner,omitempty"`
	Biografy  string             `bson:"biography,omitempty" json:"biography,omitempty"`
	Location  string             `bson:"location,omitempty" json:"location,omitempty"`
	WebSite   string             `bson:"webSite,omitempty" json:"webSite,omitempty"`
}

func (user *User) ValidateUserCreation() error {
	if user == nil {
		return errors.New("User is required")
	}

	if len(user.Email) == 0 {
		return errors.New("Email is required")
	}

	if len(user.Password) < 6 {
		return errors.New("Password have to be more than six character")
	}
	return nil
}
