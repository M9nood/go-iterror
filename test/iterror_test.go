package test

import (
	"testing"

	iterror "github.com/M9nood/go-iterror"

	"github.com/stretchr/testify/assert"
)

func runItErrors() []iterror.ErrorException {
	var errors [8]iterror.ErrorException
	errors[0] = iterror.NewError(iterror.BadRequestException, "Invalid request")
	errors[1] = iterror.NewError(iterror.UnauthorizedException, "Unauthorized")
	errors[2] = iterror.NewError(iterror.ForbiddenException, "Forbidden")
	errors[3] = iterror.NewError(iterror.NotFoundException, "Not Found")
	errors[4] = iterror.NewError(iterror.RequestTimeoutException, "Request Timeout")
	errors[5] = iterror.NewError(iterror.ConflictException, "Conflict")
	errors[6] = iterror.NewError(iterror.ManyRequestException, "Many Request")
	errors[7] = iterror.NewError(iterror.InternalServerErrorException, "Internal Server Error")

	return errors[:]
}

func TestErrors(t *testing.T) {
	errs := runItErrors()
	for _, err := range errs {
		assert.NotEmpty(t, err.GetCode())
		switch err.GetCode() {
		case "400":
			assert.Equal(t, "400", err.GetCode())
			assert.Equal(t, 400, err.GetHttpCode())
			assert.Equal(t, "BadRequestException", err.GetName())
			break
		case "401":
			assert.Equal(t, "401", err.GetCode())
			assert.Equal(t, 401, err.GetHttpCode())
			assert.Equal(t, "UnauthorizedException", err.GetName())
			break
		case "403":
			assert.Equal(t, "403", err.GetCode())
			assert.Equal(t, 403, err.GetHttpCode())
			assert.Equal(t, "ForbiddenException", err.GetName())
			break
		case "404":
			assert.Equal(t, "404", err.GetCode())
			assert.Equal(t, 404, err.GetHttpCode())
			assert.Equal(t, "NotFoundException", err.GetName())
			break
		case "408":
			assert.Equal(t, "408", err.GetCode())
			assert.Equal(t, 408, err.GetHttpCode())
			assert.Equal(t, "RequestTimeoutException", err.GetName())
			break
		case "409":
			assert.Equal(t, "409", err.GetCode())
			assert.Equal(t, 409, err.GetHttpCode())
			assert.Equal(t, "ConflictException", err.GetName())
			break
		case "429":
			assert.Equal(t, "429", err.GetCode())
			assert.Equal(t, 429, err.GetHttpCode())
			assert.Equal(t, "ManyRequestException", err.GetName())
			break
		case "500":
			assert.Equal(t, "500", err.GetCode())
			assert.Equal(t, 500, err.GetHttpCode())
			assert.Equal(t, "InternalServerErrorException", err.GetName())
			break
		}
	}
}
