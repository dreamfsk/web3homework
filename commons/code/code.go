package code

import (
	"com.dreamfsk/blog/commons"
	"com.dreamfsk/blog/config"
	_ "embed"
)

type ErrorCode int

const (
	ServerError           ErrorCode = 10101
	TooManyRequests       ErrorCode = 10102
	ParamBindError        ErrorCode = 10103
	AuthorizationError    ErrorCode = 101041
	AuthorizationInvaild  ErrorCode = 101042
	AuthorizationNotFound ErrorCode = 101043
	AuthorizationInvalid  ErrorCode = 101044
	UrlSignError          ErrorCode = 10105
	CacheSetError         ErrorCode = 10106
	CacheGetError         ErrorCode = 10107
	CacheDelError         ErrorCode = 10108
	CacheNotExist         ErrorCode = 10109
	ResubmitError         ErrorCode = 10110
	HashIdsEncodeError    ErrorCode = 10111
	HashIdsDecodeError    ErrorCode = 10112
	RBACError             ErrorCode = 10113
	RedisConnectError     ErrorCode = 10114
	MySQLConnectError     ErrorCode = 10115
	WriteConfigError      ErrorCode = 10116
	SendEmailError        ErrorCode = 10117
	MySQLExecError        ErrorCode = 10118
	GoVersionError        ErrorCode = 10119
	SocketConnectError    ErrorCode = 10120
	SocketSendError       ErrorCode = 10121

	UserCreateError             ErrorCode = 20201
	UserListError               ErrorCode = 20202
	UserDeleteError             ErrorCode = 20203
	UserUpdateError             ErrorCode = 20204
	UserResetPasswordError      ErrorCode = 20205
	UserLoginError              ErrorCode = 20206
	UserLogOutError             ErrorCode = 20207
	UserModifyPasswordError     ErrorCode = 20208
	UserModifyPersonalInfoError ErrorCode = 20209
	UserOfflineError            ErrorCode = 20212
	UserDetailError             ErrorCode = 20213

	PostCreateError ErrorCode = 20301
	PostUpdateError ErrorCode = 20302
	PostListError   ErrorCode = 20303
	PostDeleteError ErrorCode = 20304
	PostDetailError ErrorCode = 20305

	CommentCreateError ErrorCode = 20401
	CommentListError   ErrorCode = 20402
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
