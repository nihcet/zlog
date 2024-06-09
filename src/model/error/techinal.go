package error_model

import (
	"fmt"

	error_constant "github.com/nihcet/zlog/src/constant/error"
)

type TechnicalError struct {
	StatusCode  int    `json:"status_code"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
}

func (e *TechnicalError) Error() string {
	return fmt.Sprintf(string(error_constant.ErrorPattern_TechnicalError), e.StatusCode, e.ErrorCode, e.Description)
}

func (e *TechnicalError) GetStatusCode() int {
	if e.StatusCode == 0 {
		return int(error_constant.DefaultStatusCode_TechincalError)
	}

	return e.StatusCode
}

func (e *TechnicalError) GetErrorCode() string {
	if e.ErrorCode == "" {
		return string(error_constant.DefaultErrorCode_TechincalError)
	}

	return e.ErrorCode
}
