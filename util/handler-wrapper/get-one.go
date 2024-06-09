package handlerwrapper_util

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihcet/zlog/src/model"
	error_model "github.com/nihcet/zlog/src/model/error"
	"github.com/nihcet/zlog/util"
)

// expect to get non-nil response,
// else return not-found with http status 200
func GetOne[Request any, Response any](h func(context.Context, Request) (*Response, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		var req Request
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, error_model.InvalidRequestError)
			return
		}

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, error_model.InvalidRequestError)
			return
		}

		var wrappedResponse model.ResponseWrapper
		res, err := h(c, req)
		if err != nil {
			statusCode, errorDescription := util.ExtractError(err)
			c.JSON(statusCode, wrappedResponse.WithError(errorDescription))
			return
		}

		if res == nil {
			_, errorDescription := util.ExtractError(error_model.NotFoundError)
			c.JSON(http.StatusOK, wrappedResponse.WithSuccessAndErrorDetail(errorDescription))
			return
		}

		c.JSON(http.StatusOK, wrappedResponse.WithResponse(res))
	}
}
