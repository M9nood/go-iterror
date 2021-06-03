package iterror

func ErrorInternalServer(msg string) *ErrorInfo {
	err := errors[InternalServerErrorException]
	err.ErrorMessage = msg
	return &err
}
