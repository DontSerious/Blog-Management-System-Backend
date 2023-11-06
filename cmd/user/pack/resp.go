package pack

import "Bishe/be/kitex_gen/user"

func BuildBaseResponse(statusCode int64, statusMsg string) *user.BaseResponse {
	return &user.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
	}
}