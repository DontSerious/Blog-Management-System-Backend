package db

import (
	"context"
	"fmt"

	c "github.com/DontSerious/Blog-Management-System-Backend/pkg/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MyDB               *mongo.Database
	UsersCollection    *mongo.Collection
	UserInfoCollection *mongo.Collection
)

func Init() {
	clientOptions := options.Client().ApplyURI(uri())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	MyDB = client.Database(c.MONGODB_DATABASE)

	// 在这里创建 users 和 userInfo 集合
	createCollections()
	UsersCollection = MyDB.Collection("users")
	UserInfoCollection = MyDB.Collection("userInfo")
}

func uri() string {
	// 确保在admin数据库新建用户用于认证
	const format = "mongodb://%s:%s@%s:%d/%s?authSource=admin&authMechanism=SCRAM-SHA-1"
	return fmt.Sprintf(format, c.MONGODB_USER, c.MONGODB_PASSWORD, c.MONGODB_HOST_NAME, c.MONGODB_PORT, c.MONGODB_DATABASE)
}

func createCollections() {
	collections, err := MyDB.ListCollectionNames(context.TODO(), nil)
	if err != nil && err.Error() != "document is nil" {
		panic(err)
	}

	// 检查 users 是否存在
	usersExist := false
	for _, name := range collections {
		if name == "users" {
			usersExist = true
			break
		}
	}

	// 如果 users 不存在，创建它
	if !usersExist {
		err := MyDB.CreateCollection(context.TODO(), "users")
		if err != nil {
			panic(err)
		}
	}

	// 类似地，检查 userInfo 是否存在，如果不存在，创建它
	userInfoExist := false
	for _, name := range collections {
		if name == "userInfo" {
			userInfoExist = true
			break
		}
	}

	if !userInfoExist {
		err := MyDB.CreateCollection(context.TODO(), "userInfo")
		if err != nil {
			panic(err)
		}
	}
}
