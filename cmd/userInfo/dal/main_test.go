package dal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DontSerious/Blog-Management-System-Backend/cmd/userInfo/dal/db"
)

var ctx context.Context
var cancel context.CancelFunc

func TestMain(m *testing.M) {
	ctx, cancel = context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	Init()
	m.Run()
}

func TestUpdate(t *testing.T) {
	err := db.UpdateUserInfo(ctx, "655841eb9016429abbc94b6c", &db.UserInfo{
		Categories: []string{"aa", "bb", "cc"},
		Tags:       []string{"cc", "dd", "ee"},
	})
	fmt.Printf("err: %v\n", err)
}

func TestQueryUserInfo(t *testing.T) {
	res, _ := db.QueryUserInfo(ctx, "655841eb9016429abbc94b6c")
	fmt.Printf("res: %v\n", res.Categories)
}
