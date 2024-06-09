package util

import (
	"net/http"

	error_model "github.com/nihcet/zlog/src/model/error"
)

func ExtractError(e error) (int, error_model.ErrorDescription) {
	var (
		statusCode  = http.StatusInternalServerError
		errorCode   = "unknown error code"
		description = "unknown description"
	)

	switch _e := e.(type) {
	case *error_model.BusinessError:
		description = _e.Description
		statusCode = _e.GetStatusCode()
		errorCode = _e.GetErrorCode()

	case *error_model.TechnicalError:
		description = _e.Description
		statusCode = _e.GetStatusCode()
		errorCode = _e.GetErrorCode()
	}

	return statusCode, error_model.ErrorDescription{
		ErrorCode:   errorCode,
		Description: description,
	}
}
