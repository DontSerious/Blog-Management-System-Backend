package pack

import (
	"Bishe/be/cmd/userInfo/dal/db"
	userinfo "Bishe/be/kitex_gen/userInfo"
)

func UserInfo(info *db.UserInfo) *userinfo.UserInfo {
	return &userinfo.UserInfo{
		Categories: info.Categories,
		Tags: info.Tags,
	}
}