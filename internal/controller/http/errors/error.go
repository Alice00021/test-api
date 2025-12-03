package errors

import (
	"errors"
	"github.com/Alice00021/test_api/internal/entity"
	httpError "github.com/Alice00021/test_common/pkg/httpserver"
	rmqRpcError "github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(c *gin.Context, err error) {
	var httpErr httpError.HttpError
	if errors.As(err, &httpErr) {
		c.AbortWithStatusJSON(httpErr.Status, httpErr)
		return
	}

	if errors.Is(err, entity.ErrAccessDenied) {
		httpErr = httpError.NewForbiddenError(err.Error())
		c.AbortWithStatusJSON(httpErr.Status, httpErr)
		return
	}

	var rmqRpcErr *rmqRpcError.MessageError
	if errors.As(err, &rmqRpcErr) {
		switch rmqRpcErr.Code {
		case rmqRpcError.InvalidArgument:
			httpErr = httpError.NewBadRequestError(rmqRpcErr.Message)
		case rmqRpcError.Unauthorized:
			httpErr = httpError.NewUnauthorizedError(rmqRpcErr.Message)
		case rmqRpcError.Forbidden:
			httpErr = httpError.NewForbiddenError(rmqRpcErr.Message)
		case rmqRpcError.NotFound:
			httpErr = httpError.NewNotFoundError(rmqRpcErr.Message)
		case rmqRpcError.AlreadyExists:
			httpErr = httpError.NewConflictError(rmqRpcErr.Message)
		default:
			httpErr = httpError.NewInternalServerError(rmqRpcErr.Message)
		}

		c.AbortWithStatusJSON(httpErr.Status, httpErr)
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, httpError.NewInternalServerError(err))
}
