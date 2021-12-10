package response

import (
	"fmt"
)

var (
	// 鉴权类 10000
	IncorrectUsernamePassword = NewError(10000, "用户名或密码错误！")
	AccountLocked             = NewError(10001, "该账户已经被锁定，请联系管理员解锁！")
	FailedGenerateToken       = NewError(10002, "生成Token失败！")
	TokenExpire               = NewError(10003, "Token过期，请重新登录！")
	TokenError                = NewError(10004, "Token解析错误！")
	TokenNotExist             = NewError(10005, "token not exists, please login first. ")
	Forbidden                 = NewError(10006, "没有权限，禁止访问！")
	AuthenticationFailed      = NewError(10007, "认证失败！")
	EmailExisted              = NewError(10008, "该邮箱已经存在！")
	PhoneExisted              = NewError(10009, "该手机号已经注册！")
	NicknameExisted           = NewError(10010, "该昵称已经存在！")
	UsernameHasExisted        = NewError(10011, "该用户已经存在！")
	NoAddAuthorithy           = NewError(10012, "没有添加权限")
	UpdatePasswordError       = NewError(10013, "修改密码失败！")
	NotLogin                  = NewError(10014, "请先登录！")
	AuthorizationFailed       = NewError(10015, "鉴权失败！")
	// http 20000
	InvalidParams = NewError(20000, "参数错误！")

	// 文件上传  30000
	UploadFailed         = NewError(30000, "上传失败！")
	NotSupportedSuffix   = NewError(30001, "不支持该文件类型！")
	ExceededMaximumLimit = NewError(30002, "文件太大了！")
	FileNotFound         = NewError(30003, "文件不存在！")
	FailedRemoveFile     = NewError(30004, "删除存储文件失败！")

	// 数据库 40000
	DatabaseDmlError    = NewError(40000, "数据库操作错误")
	DatabaseInsertError = NewError(40001, "数据库插入错误")
	DatabaseSelectError = NewError(40002, "数据库查询错误")
	DatabaseUpdateError = NewError(40003, "数据库更新错误")
	DatabaseDeleteError = NewError(40004, "数据库删除错误")

	// 系统错误
	InternalServerError = NewError(50000, "系统内部错误")
	RecordNotFound      = NewError(50001, "该记录不存在！")
	RecoreExisted       = NewError(50002, "该记录已经存在！")

)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("{\"code\": %d, \"message\": %s}", e.Code, e.Message)
}

func (e Error) SetMsg(format string, args ...interface{}) Error {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

var codes = map[int]string{}

func NewError(code int, msg string) Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("Error code %d  has existed, please change one!", code))
	}
	codes[code] = msg
	return Error{Code: code, Message: msg}
}
