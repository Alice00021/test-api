package http

import (
	"github.com/Alice00021/test_api/config"
	"github.com/Alice00021/test_api/internal/controller/http/middleware"
	v1 "github.com/Alice00021/test_api/internal/controller/http/v1"
	"github.com/Alice00021/test_api/internal/di"
	"github.com/Alice00021/test_common/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/websocket"
	"net/http"

	_ "github.com/Alice00021/test_common/pkg/httpserver"
)

// NewRouter -.
// Swagger spec:
// @title Finance API
// @version 1.0
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func NewRouter(handler *gin.Engine, cfg *config.Config, l logger.Interface, uc *di.UseCase, ws *websocket.Server) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	if cfg.Metrics.Enabled {
		handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	// Swagger
	if cfg.Swagger.Enabled {
		handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Routers
	//publicV1Group := handler.Group("/v1")
	//{
	//	v1.NewAuthRoutes(publicV1Group, l, uc.Auth)
	//}

	privateV1Group := handler.Group("/v1")
	privateV1Group.Use(middleware.JwtAuthMiddleware(uc.Auth))
	{
		v1.NewAuthorRoutes(privateV1Group, l, uc.Author)
		v1.NewBookRoutes(privateV1Group, l, uc.Book)
		v1.NewCommandRoutes(privateV1Group, l, uc.Command)
		v1.NewOperationRoutes(privateV1Group, l, uc.Operation)
	}

}
