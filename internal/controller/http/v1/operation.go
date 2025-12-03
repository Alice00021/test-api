package v1

import (
	"github.com/Alice00021/test_api/internal/controller/http/errors"
	"github.com/Alice00021/test_api/internal/controller/http/middleware"
	"github.com/Alice00021/test_api/internal/controller/http/v1/request"
	auth "github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_api/internal/usecase"
	"github.com/Alice00021/test_api/internal/utils"
	httpError "github.com/Alice00021/test_common/pkg/httpserver"
	"github.com/Alice00021/test_common/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type operationRoutes struct {
	l  logger.Interface
	uc usecase.Operation
}

func NewOperationRoutes(privateGroup *gin.RouterGroup, l logger.Interface, uc usecase.Operation) {
	r := &operationRoutes{l, uc}

	{
		h := privateGroup.Group("/operations").Use(
			middleware.IsRolesMiddleware(auth.UserRoleAdmin, auth.UserRoleClient))
		h.POST("", r.createOperation)
		h.PUT("/:id", r.updateOperation)
		h.GET("", r.getOperations)
		h.GET("/:id", r.getOperation)
		h.DELETE("/:id", r.deleteOperation)

	}
}

func (r *operationRoutes) getOperations(c *gin.Context) {
	res, err := r.uc.GetOperations(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - getOperations")
		errors.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (r *operationRoutes) getOperation(c *gin.Context) {
	id, err := utils.ParsePathParam(utils.ParseParams{Context: c, Key: "id"}, utils.ParseInt64)
	if err != nil {
		r.l.Error(err, "http - v1 - getOperation")
		errors.ErrorResponse(c, httpError.NewBadPathParamsError(err))
		return
	}

	res, err := r.uc.GetOperation(c, id)
	if err != nil {
		r.l.Error(err, "http - v1 - getOperation")
		errors.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (r *operationRoutes) createOperation(c *gin.Context) {
	var req request.CreateOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "http - v1 - createOperation")
		errors.ErrorResponse(c, httpError.NewBadRequestBodyError(err))
		return
	}

	inp := req.ToEntity()

	res, err := r.uc.CreateOperation(c.Request.Context(), inp)
	if err != nil {
		r.l.Error(err, "http - v1 - createOperation")
		errors.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (r *operationRoutes) updateOperation(c *gin.Context) {
	id, err := utils.ParsePathParam(utils.ParseParams{Context: c, Key: "id"}, utils.ParseInt64)
	if err != nil {
		r.l.Error(err, "http - v1 - updateOperation")
		errors.ErrorResponse(c, httpError.NewBadPathParamsError(err))
		return
	}

	var req request.UpdateOperationRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "http - v1 - updateOperation")
		errors.ErrorResponse(c, httpError.NewBadRequestBodyError(err))
		return
	}

	inp := req.ToEntity()
	inp.ID = id

	if err = r.uc.UpdateOperation(c.Request.Context(), inp); err != nil {
		r.l.Error(err, "http - v1 - updateOperation")
		errors.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *operationRoutes) deleteOperation(c *gin.Context) {
	id, err := utils.ParsePathParam(utils.ParseParams{Context: c, Key: "id"}, utils.ParseInt64)
	if err != nil {
		r.l.Error(err, "http - v1 - deleteOperation")
		errors.ErrorResponse(c, httpError.NewBadPathParamsError(err))
		return
	}

	err = r.uc.DeleteOperation(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - deleteOperation")
		errors.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}
