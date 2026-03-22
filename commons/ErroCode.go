package commons

type ErrorCode int

const (
	ApiErr ErrorCode = iota + 1000
	Unauthorized
	UserAddError
	UserNotFound
	UserExist
	UserEmailExist
	UserInvalid
	UserUpdateFail
	PostAddError
	PostListError
	PostNotFoundError
	PostUpdateError
	PostUpdateAuthError
	PostDeleteAuthError
	PostDeleteError
	CommentCreateError
)

func (s ErrorCode) String() string {
	if s < 1000 || s > 1006 {
		return "Unknown error"
	}
	return [...]string{
		"Internal server error",
		"Unauthorized",
		"User add fail",
		"User not found",
		"User already exists",
		"User email already exists",
		"Invalid credentials",
		"User update fail",
		"post add fail",
		"post list search fail",
		"post not found fail",
		"post update content fail",
		"only author can update the post content",
		"only author can delete the post",
		"post delete fail",
		"comment add fail",
	}[s-1000]
}

// 添加额外的方法，提高可用性
func (s ErrorCode) Code() int {
	return int(s)
}

func (s ErrorCode) Error() string {
	return s.String()
}
