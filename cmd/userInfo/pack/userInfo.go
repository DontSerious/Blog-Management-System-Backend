package pack

import (
	"github.com/DontSerious/Blog-Management-System-Backend/cmd/userInfo/dal/db"
	userinfo "github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/userInfo"
)

func UserInfo(info *db.UserInfo) *userinfo.UserInfo {
	return &userinfo.UserInfo{
		Categories: info.Categories,
		Tags:       info.Tags,
	}
}
