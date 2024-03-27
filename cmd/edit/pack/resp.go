package pack

import "github.com/DontSerious/Blog-Management-System-Backend/kitex_gen/edit"

func BuildBaseResponse(statusCode int64, statusMsg string) *edit.BaseResponse {
	return &edit.BaseResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
