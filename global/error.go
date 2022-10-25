package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	ValidateError CustomError
	ServiceError  CustomError
}

var Errors = CustomErrors{
	ValidateError: CustomError{40001, "请求参数错误"},
	ServiceError:  CustomError{50001, "service错误"},
}
