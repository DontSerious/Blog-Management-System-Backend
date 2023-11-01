package dal

import (
	"context"
	"testing"
	"time"

	"github.com/DontSerious/DouSheng/cmd/user/dal/db"
)

var ctx context.Context
var cancel context.CancelFunc

func TestMain(m *testing.M) {
	ctx, cancel = context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	Init()
	m.Run()
}

func TestQueryUserList(t *testing.T) {
	//将记录写入数据库
	db.CreateUser(ctx, []*db.User{{
		Name:          "444",
		Password:      "444",
		FollowCount:   0,
		FollowerCount: 0,
	}})
}

func TestFollowSB(t *testing.T) {
	db.Follow(ctx, 2, 1, 2)
}