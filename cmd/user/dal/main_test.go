package dal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"Bishe/be/cmd/user/dal/db"
)

var ctx context.Context
var cancel context.CancelFunc

func TestMain(m *testing.M) {
	ctx, cancel = context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	Init()
	m.Run()
}

func TestCreate(t *testing.T) {
	db.CreateUser(ctx, &db.User{
		UserName: "u1",
		Password: "p1",
	})
}

func TestQueryUser(t *testing.T) {
	user, _ := db.QueryUser(ctx, "qq")
	fmt.Printf("user.ObjectID.Hex(): %v\n", user.ObjectID.Hex())
}