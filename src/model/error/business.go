package error_model

import (
	"fmt"

	error_constant "github.com/nihcet/zlog/src/constant/error"
)

type BusinessError struct {
	StatusCode  int    `json:"status_code"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf(string(error_constant.ErrorPattern_BusinessError), e.StatusCode, e.ErrorCode, e.Description)
}

func (e *BusinessError) GetStatusCode() int {
	if e.StatusCode == 0 {
		return int(error_constant.DefaultStatusCode_BusinessError)
	}

	return e.StatusCode
}

func (e *BusinessError) GetErrorCode() string {
	if e.ErrorCode == "" {
		return string(error_constant.DefaultErrorCode_BusinessError)
	}

	return e.ErrorCode
}
