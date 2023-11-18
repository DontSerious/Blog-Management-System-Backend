package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ObjectID 	primitive.ObjectID 	`bson:"_id" json:"_id"`
	Categories	[]string			`bson:"categories" json:"categories"`
 	Tags		[]string			`bson:"tags" json:"tags"`
}

func UpdateUserInfo(ctx context.Context, userId string, userInfo *UserInfo) error {
	// convert
	id, err := primitive.ObjectIDFromHex(userId)
    if err != nil {
        return err
    }

	// userInfo表插入
	_, err = UserInfoCollection.UpdateByID(
		context.TODO(), 
		id,
		bson.D{
			{
				Key: "$set", Value:  bson.D{
				{Key: "categories", Value: userInfo.Categories},
				{Key: "tags", Value: userInfo.Tags},
			}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func QueryUserInfo(ctx context.Context, userId string) (*UserInfo, error) {
	var userInfo UserInfo

	// convert
	id, err := primitive.ObjectIDFromHex(userId)
    if err != nil {
        return &userInfo, err
    }

	res := UserInfoCollection.FindOne(ctx, bson.M{"_id": id})
	if err := res.Decode(&userInfo); err != nil {
		return &userInfo, err
	}

	return &userInfo, nil
}