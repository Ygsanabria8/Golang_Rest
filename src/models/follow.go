package models

type Follow struct {
	UserId       string `bson:"userId,omitempty" json:"userId"`
	UserFollowId string `bson:"userFollowId,omitempty" json:"userFollowId"`
	IsFollowed   bool   `bson:"isFollowed,omitempty" json:"isFollowed"`
}
