package error_model

import error_constant "github.com/nihcet/zlog/src/constant/error"

type InternalError interface {
	Error() string

	GetStatusCode() int
	GetErrorCode() string
}

var (
	InvalidRequestError = &TechnicalError{
		ErrorCode:   string(error_constant.InvalidRequest_ErrorCode),
		Description: "unable to parse data from request",
	}

	NotFoundError = &BusinessError{
		ErrorCode:   string(error_constant.NotFound_ErrorCode),
		Description: "not found data for the specify detail(s)",
	}
)

type ErrorDescription struct {
	Description string
	ErrorCode   string
}
