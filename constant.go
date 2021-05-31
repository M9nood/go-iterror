package iterror

var (
	BadRequestException          = "400"
	UnauthorizedException        = "401"
	ForbiddenException           = "403"
	NotFoundException            = "404"
	RequestTimeoutException      = "408"
	ConflictException            = "409"
	ManyRequestException         = "429"
	InternalServerErrorException = "500"
)

var errors = map[string]ErrorInfo{
	"400": ErrorInfo{ErrorCode: "400", ErrorHttpCode: 400, ErrorName: "BadRequestException"},
	"401": ErrorInfo{ErrorCode: "401", ErrorHttpCode: 401, ErrorName: "UnauthorizedException"},
	"403": ErrorInfo{ErrorCode: "403", ErrorHttpCode: 403, ErrorName: "ForbiddenException"},
	"404": ErrorInfo{ErrorCode: "404", ErrorHttpCode: 404, ErrorName: "NotFoundException"},
	"408": ErrorInfo{ErrorCode: "408", ErrorHttpCode: 408, ErrorName: "RequestTimeoutException"},
	"409": ErrorInfo{ErrorCode: "409", ErrorHttpCode: 409, ErrorName: "ConflictException"},
	"429": ErrorInfo{ErrorCode: "429", ErrorHttpCode: 429, ErrorName: "ManyRequestException"},
	"500": ErrorInfo{ErrorCode: "500", ErrorHttpCode: 500, ErrorName: "InternalServerErrorException"},
}
