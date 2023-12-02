package transport

import (
	"net/http"

	"github.com/dmytrodemianchuk/go-crud-csv/internal/config"
	"github.com/dmytrodemianchuk/go-crud-csv/internal/services"
	v1 "github.com/dmytrodemianchuk/go-crud-csv/internal/transport/v1"

	_ "github.com/dmytrodemianchuk/go-crud-csv/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
