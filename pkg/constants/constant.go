package constants

// MongoDB config
const (
	MONGODB_DATABASE  = "bishe"
	MONGODB_USER      = "bishe"
	MONGODB_PASSWORD  = "bishe"
	MONGODB_HOST_NAME = "8.222.170.236"
	MONGODB_PORT      = 27777
)

// addr
const (
	EtcdAddress            = "127.0.0.1:2379"
	UserResolveTCPAddr     = "127.0.0.1:6660"
	UserInfoResolveTCPAddr = "127.0.0.1:6661"
	EditResolveTCPAddr     = "127.0.0.1:6662"

	UserServiceName     = "user"
	UserInfoServiceName = "userInfo"
	EditServiceName     = "edit"
)

// editService
const (
	EditDirectory = "/root/blog"
)
