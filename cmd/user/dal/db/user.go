package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ObjectID 	primitive.ObjectID 	`bson:"_id" json:"_id"`
	UserName	string				`bson:"username" json:"username"`
	Password	string				`bson:"password" json:"password"`
}

func CreateUser(ctx context.Context, user *User) (string, error) {
	doc := bson.D{
		{Key: "username", Value: user.UserName},
		{Key: "password", Value: user.Password},
	}

	// users表插入
	res, err := UsersCollection.InsertOne(context.TODO(), doc)
    if err != nil {
        return "", err
    }

	// 获取 _id
	insId := res.InsertedID.(primitive.ObjectID)
	idStr := insId.Hex()
	doc = bson.D{
		{Key: "_id", Value: insId},
		{Key: "username", Value: user.UserName},
	}

	// userInfo表插入
	_, err = UserInfoCollection.InsertOne(context.TODO(), doc)
    if err != nil {
        return "", err
    }

	return idStr, nil
}

func QueryUser(ctx context.Context, userName string) (User, error) {
	res := UsersCollection.FindOne(ctx, bson.M{"username": userName})

	var user User
	if err := res.Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}
