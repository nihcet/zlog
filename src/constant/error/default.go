package error_constant

import (
	"net/http"
)

type ErrorPattern string

const (
	ErrorPattern_BusinessError  ErrorPattern = "[%d] %s - %+s"
	ErrorPattern_TechnicalError ErrorPattern = "[%d] %s - %+s"
)

type DefaultStatusCode int

const (
	DefaultStatusCode_BusinessError  DefaultStatusCode = http.StatusInternalServerError
	DefaultStatusCode_TechincalError DefaultStatusCode = http.StatusInternalServerError
)

type DefaultErrorCode string

const (
	DefaultErrorCode_BusinessError  DefaultErrorCode = "BIZN-0000"
	DefaultErrorCode_TechincalError DefaultErrorCode = "TECH-0000"
)
