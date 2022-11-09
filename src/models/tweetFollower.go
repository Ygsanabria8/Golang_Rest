package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TweetFollower struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserId       string             `bson:"userId" json:"userId,omitempty"`
	UserFollowId string             `bson:"userFollowId" json:"userFollowId,omitempty"`
	Tweet        Tweet
}
