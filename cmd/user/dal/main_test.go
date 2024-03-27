package dal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/dal/db"
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

func TestGetAllUser(t *testing.T) {
	users, _ := db.GetAllUser(ctx)
	for _, user := range users {
		fmt.Print(user)
	}
}

func TestDelUser(t *testing.T) {
	err := db.DelUser(ctx, "658e2e9aee48b155594a9584")
	if err != nil {
		fmt.Print(err)
	}
}

func TestChangePWD(t *testing.T) {
	err := db.ChangePWD(ctx, "658e2e7cee48b155594a9582", "qw")
	if err != nil {
		fmt.Print(err)
	}
}
