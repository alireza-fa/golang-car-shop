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
}