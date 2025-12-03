package v1

import (
	"github.com/Alice00021/test_api/internal/controller/http/errors"
	"github.com/Alice00021/test_api/internal/controller/http/middleware"
	auth "github.com/Alice00021/test_api/internal/entity/back"
	"github.com/Alice00021/test_api/internal/usecase"
	"github.com/Alice00021/test_common/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type commandRoutes struct {
	l  logger.Interface
	uc usecase.Command
}

func NewCommandRoutes(privateGroup *gin.RouterGroup, l logger.Interface, uc usecase.Command) {
	r := &commandRoutes{l, uc}

	{
		h := privateGroup.Group("/commands").Use(
			middleware.IsRolesMiddleware(auth.UserRoleAdmin, auth.UserRoleClient))
		h.POST("", r.updateCommands)
		h.GET("", r.getCommands)
	}
}

func (r *commandRoutes) getCommands(c *gin.Context) {
	res, err := r.uc.GetCommands(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - getCommands")
		errors.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (r *commandRoutes) updateCommands(c *gin.Context) {
	if err := r.uc.UpdateCommands(c.Request.Context()); err != nil {
		r.l.Error(err, "http - v1 - updateCommands")
		errors.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}
