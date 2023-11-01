package constants

const (
	UserTableName          = "users"
	UserServiceName        = "user"
	FavoriteTableName      = "favorites"
	FavoriteServiceName    = "favorite"
	RelationTableName      = "relations"
	RelationServiceName    = "relation"
	VideoTableName         = "videos"
	VideoServiceName       = "video"
	CommentTableName       = "comments"
	CommentServiceName     = "comment"
	MySQLTestDSN           = "joker:!joker2233@tcp(rm-7xv18y05c877450825o.mysql.rds.aliyuncs.com:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddress            = "127.0.0.1:2379"
	UserResolveTCPAddr     = "127.0.0.1:6660"
	RelationResolveTCPAddr = "127.0.0.1:6661"
	FavoriteResolveTCPAddr = "127.0.0.1:6602"
	CommentResolveTCPAddr  = "127.0.0.1:6603"
	VideoResolveTCPAddr    = "127.0.0.1:6604"
)
