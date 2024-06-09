package handlerwrapper_util

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihcet/zlog/src/model"
	error_model "github.com/nihcet/zlog/src/model/error"
	"github.com/nihcet/zlog/util"
)

func Post[Request any, Response any](h func(context.Context, Request) (*Response, error)) func(*gin.Context) {
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
			statusCode, errorResponse := util.ExtractError(err)
			c.JSON(statusCode, wrappedResponse.WithError(errorResponse))
			return
		}

		if res == nil {
			c.JSON(http.StatusCreated, wrappedResponse.WithCreateSuccessResponse())
			return
		}

		c.JSON(http.StatusOK, wrappedResponse.WithResponse(res))
	}
}
