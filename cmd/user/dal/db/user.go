package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ObjectID   primitive.ObjectID `bson:"_id" json:"_id"`
	UserName   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"password"`
	Categories []string           `bson:"categories" json:"categories"`
	Tags       []string           `bson:"tags" json:"tags"`
}

func CreateUser(ctx context.Context, user *User) (string, error) {
	// users表插入
	res, err := UsersCollection.InsertOne(context.TODO(), &bson.D{
		{Key: "username", Value: user.UserName},
		{Key: "password", Value: user.Password},
	})
	if err != nil {
		return "", err
	}

	// 获取 _id，userInfo表插入
	insId := res.InsertedID.(primitive.ObjectID)
	idStr := insId.Hex()
	_, err = UserInfoCollection.InsertOne(context.TODO(), &bson.D{
		{Key: "_id", Value: insId},
	})
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

func ChangePWD(ctx context.Context, userId string, newPassword string) error {
	// convert
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	// 构建更新条件
	filter := bson.M{"_id": id}
	// 构建更新内容
	update := bson.M{"$set": bson.M{"password": newPassword}}

	// 执行更新操作
	_, err = UsersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DelUser(ctx context.Context, userId string) error {
	// convert
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	// 构建条件
	filter := bson.M{"_id": id}

	// 执行删除操作
	_, err = UsersCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	_, err = UserInfoCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUser(ctx context.Context) ([]*User, error) {
	// 构建空的筛选条件，以获取集合中的所有文档
	filter := bson.M{}

	// 执行查询操作
	usersCursor, err := UsersCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer usersCursor.Close(ctx)

	userInfoCursor, err := UserInfoCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer userInfoCursor.Close(ctx)

	// 遍历结果集
	var users []*User
	for usersCursor.Next(ctx) && userInfoCursor.Next(ctx) {
		var user User

		err := usersCursor.Decode(&user)
		if err != nil {
			return nil, err
		}

		err = userInfoCursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
