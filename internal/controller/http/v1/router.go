package v1

import (
	"github.com/monferon/apichka/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/monferon/apichka/docs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Translation) {

// NewRouter -.
// Swagger spec:
// @title       Go Apichka
// @description Some service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//Routers
	h := handler.Group("/v1")
	{
		newTestRouter(h, l)
	}
}
