package db

import (
	c "Bishe/be/pkg/constants"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    MyDB *mongo.Database
    UserInfoCollection *mongo.Collection
)

func Init() {
	clientOptions := options.Client().ApplyURI(uri())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	MyDB = client.Database(c.MONGODB_DATABASE)

	// 在这里创建 userInfo 集合
    createCollection()
	UserInfoCollection = MyDB.Collection("userInfo")
}

func uri() string {
    // 确保在admin数据库新建用户用于认证
	const format = "mongodb://%s:%s@%s:%d/%s?authSource=admin&authMechanism=SCRAM-SHA-1"
	return fmt.Sprintf(format, c.MONGODB_USER, c.MONGODB_PASSWORD, c.MONGODB_HOST_NAME, c.MONGODB_PORT, c.MONGODB_DATABASE)
}

func createCollection() {
    collections, err := MyDB.ListCollectionNames(context.TODO(), nil)
    if err != nil && err.Error() != "document is nil"  {
        panic(err)
    }

    // 检查 userInfo 是否存在，如果不存在，创建它
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