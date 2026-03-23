package code

import (
	"com.dreamfsk/blog/commons"
	"com.dreamfsk/blog/config"
	_ "embed"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	UrlSignError       = 10105
	CacheSetError      = 10106
	CacheGetError      = 10107
	CacheDelError      = 10108
	CacheNotExist      = 10109
	ResubmitError      = 10110
	HashIdsEncodeError = 10111
	HashIdsDecodeError = 10112
	RBACError          = 10113
	RedisConnectError  = 10114
	MySQLConnectError  = 10115
	WriteConfigError   = 10116
	SendEmailError     = 10117
	MySQLExecError     = 10118
	GoVersionError     = 10119
	SocketConnectError = 10120
	SocketSendError    = 10121

	UserCreateError             = 20201
	UserListError               = 20202
	UserDeleteError             = 20203
	UserUpdateError             = 20204
	UserResetPasswordError      = 20205
	UserLoginError              = 20206
	UserLogOutError             = 20207
	UserModifyPasswordError     = 20208
	UserModifyPersonalInfoError = 20209
	UserOfflineError            = 20212
	UserDetailError             = 20213

	PostCreateError = 20301
	PostUpdateError = 20302
	PostListError   = 20303
	PostDeleteError = 20304
	PostDetailError = 20305

	CommentCreateError = 20401
	CommentListError   = 20402
)

func Text(code int) string {
	lang := config.Env().Language.Local

	if lang == commons.ZhCN {
		return zhCNText[code]
	}

	if lang == commons.EnUS {
		return enUSText[code]
	}

	return zhCNText[code]
}
