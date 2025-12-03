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

type bookRoutes struct {
	l  logger.Interface
	uc usecase.Book
}

func NewBookRoutes(privateGroup *gin.RouterGroup, l logger.Interface, uc usecase.Book) {
	r := &bookRoutes{l, uc}

	{
		h := privateGroup.Group("/books").Use(
			middleware.IsRolesMiddleware(auth.UserRoleAdmin, auth.UserRoleClient))
		h.POST("", r.createBook)
		h.PUT("/:id", r.updateBook)
		h.GET("", r.getBooks)
		h.GET("/:id", r.getBook)
		h.DELETE("/:id", r.deleteBook)

	}
}

func (r *bookRoutes) getBooks(c *gin.Context) {
	res, err := r.uc.GetBooks(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - getBooks")
		errors.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (r *bookRoutes) getBook(c *gin.Context) {
	id, err := utils.ParsePathParam(utils.ParseParams{Context: c, Key: "id"}, utils.ParseInt64)
	if err != nil {
		r.l.Error(err, "http - v1 - getBook")
		errors.ErrorResponse(c, httpError.NewBadPathParamsError(err))
		return
	}

	res, err := r.uc.GetBook(c, id)
	if err != nil {
		r.l.Error(err, "http - v1 - getBook")
		errors.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (r *bookRoutes) createBook(c *gin.Context) {
	var req request.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "http - v1 - createBook")
		errors.ErrorResponse(c, httpError.NewBadRequestBodyError(err))
		return
	}

	inp := req.ToEntity()

	res, err := r.uc.CreateBook(c.Request.Context(), inp)
	if err != nil {
		r.l.Error(err, "http - v1 - createBook")
		errors.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (r *bookRoutes) updateBook(c *gin.Context) {
	id, err := utils.ParsePathParam(utils.ParseParams{Context: c, Key: "id"}, utils.ParseInt64)
	if err != nil {
		r.l.Error(err, "http - v1 - updateBook")
		errors.ErrorResponse(c, httpError.NewBadPathParamsError(err))
		return
	}

	var req request.UpdateBookRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "http - v1 - updateBook")
		errors.ErrorResponse(c, httpError.NewBadRequestBodyError(err))
		return
	}

	inp := req.ToEntity()
	inp.ID = id

	if err = r.uc.UpdateBook(c.Request.Context(), inp); err != nil {
		r.l.Error(err, "http - v1 - updateBook")
		errors.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (r *bookRoutes) deleteBook(c *gin.Context) {
	id, err := utils.ParsePathParam(utils.ParseParams{Context: c, Key: "id"}, utils.ParseInt64)
	if err != nil {
		r.l.Error(err, "http - v1 - deleteBook")
		errors.ErrorResponse(c, httpError.NewBadPathParamsError(err))
		return
	}

	err = r.uc.DeleteBook(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - deleteBook")
		errors.ErrorResponse(c, err)
		return
	}

	c.Status(http.StatusOK)
}
