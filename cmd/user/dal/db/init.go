package db

import (
	"Bishe/be/pkg/constants"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MyDB *mongo.Client

func Init() {
	clientOptions := options.Client().ApplyURI(uri())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	MyDB = client
}

func uri() string {
	const format = "mongodb://%s:%s@%s:%d/%s"
	return fmt.Sprintf(format, constants.MONGODB_USER, constants.MONGODB_PASSWORD, constants.MONGODB_HOST_NAME, constants.MONGODB_PORT, constants.MONGODB_DATABASE)
}