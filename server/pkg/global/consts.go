package global

const (
	Private = 1 // 私有
	Public  = 2 // 公开

	Commit  = 1 // 提交
	Staged  = 2 // 草稿
	Publish = 3 // 发布
	Delete  = 4 // 删除

	DefaultAvatar        = 1 // 默认头像图片ID
	DefaultSubjectAvatar = 1 // 默认专题头像
	DefaultSubjectCover  = 2 // 默认专题首页图片

	SessionUserIDKey     = "current_user_id"
	SessionUserNameKey   = "current_user_name"
	SessionPermissionKey = "current_user_permissions"
	SessionRoleKey       = "current_user_roles"
	SessionIsLoginKey    = "is_login"

	RequestHeaderTokenKey  = "token"
	RequestQueryTokenKey   = "token"
	RequestAccessTokenKey  = "access_token"
	RequestRefreshTokenKey = "refresh_token"
	RequestIDKey           = "requestID"
)
