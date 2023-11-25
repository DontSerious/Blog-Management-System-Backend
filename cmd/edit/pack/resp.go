package pack

import "Bishe/be/kitex_gen/edit"

func BuildBaseResponse(statusCode int64, statusMsg string) *edit.BaseResponse {
	return &edit.BaseResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
