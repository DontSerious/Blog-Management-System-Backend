package dal

import (
	"context"
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

func TestPing(t *testing.T) {
	db.Init()
}