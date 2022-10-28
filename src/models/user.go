package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User type of object for a user
type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name      string             `bson:"name, omitempty" json:"name"`
	LastName  string             `bson:"lastName, omitempty" json:"lastName"`
	DateBirth time.Time          `bson:"dateBirth, omitempty" json:"dateBirth"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biografy  string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	WebSite   string             `bson:"webSite" json:"webSite,omitempty"`
}
