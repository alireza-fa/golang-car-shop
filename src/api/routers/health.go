package routers

import (
	"github.com/alireza-fa/golang-car-shop/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.HandlerGet)
	r.POST("/", handler.HandlerPost)
	r.GET("/:id", handler.HandlerById)
}
