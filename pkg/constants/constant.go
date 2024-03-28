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
	EtcdIP                 = "etcd"
	EtcdAddress            = EtcdIP + ":2379"
	UserResolveTCPAddr     = "127.0.0.1:6660"
	UserInfoResolveTCPAddr = "127.0.0.1:6661"
	EditResolveTCPAddr     = "127.0.0.1:6662"
	FileResolveTCPAddr     = "127.0.0.1:6663"
	MainAPIPort            = ":8888"

	UserServiceName     = "user"
	UserInfoServiceName = "userInfo"
	EditServiceName     = "edit"
	FileServiceName     = "file"
)

// editService
const (
	EditDirectory = "/root/blog"
)
