package errno

import (
	"errors"
)

const (
	SuccessCode            int64 = 0
	ServiceErrCode         int64 = 10001
	ParamErrCode           int64 = 10002
	LoginErrCode           int64 = 10003
	UserNotExistErrCode    int64 = 10004
	UserNameHasUsedErrCode int64 = 10005
	TokenInvalidErrCode    int64 = 10006
	FavoriteErrCode        int64 = 10007
	WrongOperationErrCode  int64 = 10008
)

var (
	Success    			= NewErrNo(SuccessCode, "成功")
	ServiceErr 			= NewErrNo(ServiceErrCode, "服务启动失败")
	ParamErr   			= NewErrNo(ParamErrCode, "参数有误")
	LoginErr			= NewErrNo(LoginErrCode, "用户名或密码错误")
	UserNotExistErr 	= NewErrNo(UserNotExistErrCode, "用户名不存在")
	UserNameHasUsedErr  = NewErrNo(UserNameHasUsedErrCode, "用户名已存在")
	TokenErr   			= NewErrNo(TokenInvalidErrCode, "Token错误或已过期")
	WrongOperationErr   = NewErrNo(WrongOperationErrCode, "系统错误或操作不合法")
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	// return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
	return e.ErrMsg
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
