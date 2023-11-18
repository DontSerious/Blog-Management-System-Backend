package pack

import userinfo "Bishe/be/kitex_gen/userInfo"

func BuildBaseResponse(statusCode int64, statusMsg string) *userinfo.BaseResponse {
	return &userinfo.BaseResponse{
		StatusCode: statusCode,
		StatusMsg: statusMsg,
	}
}