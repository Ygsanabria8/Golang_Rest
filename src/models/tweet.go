package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId    string             `bson:"userId,omitempty" json:"userId,omitempty"`
	Message   string             `bson:"message,omitempty" json:"message,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
