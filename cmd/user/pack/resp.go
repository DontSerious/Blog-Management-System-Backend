package pack

import (
	"github.com/DontSerious/Blog-Management-System-Backend/cmd/user/dal/db"
	"github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/user"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *user.BaseResponse {
	return &user.BaseResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}

func BuildUserList(users []*db.User) []*user.User {
	var userList []*user.User

	for _, item := range users {
		user := &user.User{
			UserId:   item.ObjectID.Hex(),
			Username: item.UserName,
			UserInfo: &user.UserInfo{
				Categories: item.Categories,
				Tags:       item.Tags,
			},
		}

		userList = append(userList, user)
	}

	return userList
}
