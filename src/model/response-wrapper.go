package model

import (
	"github.com/nihcet/zlog/src/constant"
	error_model "github.com/nihcet/zlog/src/model/error"
)

type ResponseWrapper struct {
	Data        any    `json:"data"`
	IsSuccess   bool   `json:"is_success"`
	Description string `json:"description"`
	ErrorCode   string `json:"error_code"`
}

func (r ResponseWrapper) setData(response any) ResponseWrapper {
	r.Data = response
	return r
}

func (r ResponseWrapper) setIsSuccess(isSuccess bool) ResponseWrapper {
	r.IsSuccess = isSuccess
	return r
}

func (r ResponseWrapper) setDescription(description string) ResponseWrapper {
	r.Description = description
	return r
}

func (r ResponseWrapper) setErrorCode(errorCode string) ResponseWrapper {
	r.ErrorCode = errorCode
	return r
}

func (r ResponseWrapper) WithError(errorDescription error_model.ErrorDescription) ResponseWrapper {
	return r.setDescription(errorDescription.Description).setErrorCode(errorDescription.ErrorCode)
}

// similar with `WithError` but `is_success` will be 'true'
func (r ResponseWrapper) WithSuccessAndErrorDetail(errorDescription error_model.ErrorDescription) ResponseWrapper {
	return r.setIsSuccess(true).
		setDescription(errorDescription.Description).
		setErrorCode(errorDescription.ErrorCode)
}

func (r ResponseWrapper) WithResponse(response any) ResponseWrapper {
	return r.setData(response).setIsSuccess(true)
}

func (r ResponseWrapper) WithCreateSuccessResponse() ResponseWrapper {
	return r.setIsSuccess(true).setDescription(string(constant.DefaultSuccessDescription_Created))
}

func (r ResponseWrapper) WithUpdateSuccessResponse() ResponseWrapper {
	return r.setIsSuccess(true).setDescription(string(constant.DefaultSuccessDescription_Updated))
}

func (r ResponseWrapper) WithDeleteSuccessResponse() ResponseWrapper {
	return r.setIsSuccess(true).setDescription(string(constant.DefaultSuccessDescription_Deleted))
}
