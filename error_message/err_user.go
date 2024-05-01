package error_message

import "errors"

var (
	UserConflict = errors.New("Người dùng đã tồn tại")
	SignUpFail =errors.New("Đăng kí thất bại")
	UserNotFound=errors.New("Người dùng không tồn tại")

)


