package iterror

func ErrorBadRequest(msg string) *ErrorInfo {
	err := errors[BadRequestException]
	err.ErrorMessage = msg
	return &err
}

func ErrorUnauthorized(msg string) *ErrorInfo {
	err := errors[UnauthorizedException]
	err.ErrorMessage = msg
	return &err
}

func ErrorForbidden(msg string) *ErrorInfo {
	err := errors[ForbiddenException]
	err.ErrorMessage = msg
	return &err
}

func ErrorNotFound(msg string) *ErrorInfo {
	err := errors[NotFoundException]
	err.ErrorMessage = msg
	return &err
}

func ErrorRequestTimeout(msg string) *ErrorInfo {
	err := errors[RequestTimeoutException]
	err.ErrorMessage = msg
	return &err
}

func ErrorConflict(msg string) *ErrorInfo {
	err := errors[ConflictException]
	err.ErrorMessage = msg
	return &err
}

func ErrorManyRequest(msg string) *ErrorInfo {
	err := errors[ManyRequestException]
	err.ErrorMessage = msg
	return &err
}
