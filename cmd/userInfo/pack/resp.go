package pack

import userinfo "github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/userInfo"

func BuildBaseResponse(statusCode int64, statusMsg string) *userinfo.BaseResponse {
	return &userinfo.BaseResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
