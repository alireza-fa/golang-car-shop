package routers

import (
	"github.com/alireza-fa/golang-car-shop/api/handlers"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/gin-gonic/gin"
)

func CarType(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarTypeHandler(cfg)

	r.POST("/", h.Create)
	r.PATCH("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Gearbox(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewGearboxHandler(cfg)

	r.POST("/", h.Create)
	r.PATCH("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModel(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelHandler(cfg)

	r.POST("/", h.Create)
	r.PATCH("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModelColor(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelColorHandler(cfg)

	r.POST("/", h.Create)
	r.PATCH("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
