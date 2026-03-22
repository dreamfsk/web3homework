package commons

type ErrorCode int

const (
	ApiErr ErrorCode = iota + 1000
	UserAddError
	UserNotFound
	UserExist
	UserEmailExist
	UserInvalid
	UserUpdateFail
)

func (s ErrorCode) String() string {
	if s < 1000 || s > 1006 {
		return "Unknown error"
	}
	return [...]string{
		"Internal server error",
		"User add fail",
		"User not found",
		"User already exists",
		"User email already exists",
		"Invalid credentials",
		"User update fail",
	}[s-1000]
}

// 添加额外的方法，提高可用性
func (s ErrorCode) Code() int {
	return int(s)
}

func (s ErrorCode) Error() string {
	return s.String()
}
