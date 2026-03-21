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
	return [...]string{
		"Internal server error",
		"User add fail",
		"User already exists",
		"User email already exists",
		"User not found",
		"Invalid credentials",
		"User update fail",
	}[s]
}
