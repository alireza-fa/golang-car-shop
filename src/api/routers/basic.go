package routers

import (
	"github.com/alireza-fa/golang-car-shop/api/handlers"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/gin-gonic/gin"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewCountryHandler(cfg)

	r.POST("/", handler.Create)
	r.PATCH("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
	r.GET("/:id", handler.GetById)
	r.POST("/get-by-filter", handler.GetByFilter)
}

func City(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewCityHandler(cfg)

	r.POST("/", handler.Create)
	r.PATCH("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
	r.GET("/:id", handler.GetById)
	r.POST("/get-by-filter", handler.GetByFilter)
}

func File(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewFileHandler(cfg)

	r.POST("/", handler.Create)
	r.PATCH("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
	r.GET("/:id", handler.GetById)
	r.POST("/get-by-filter", handler.GetByFilter)
}

func Company(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCompanyHandler(cfg)

	r.POST("/", h.Create)
	r.PATCH("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Color(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewColorHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
