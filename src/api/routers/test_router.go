package routers

import (
	"github.com/alireza-fa/golang-car-shop/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	r.GET("/users", handler.Users)
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.UserById)
	r.GET("/users/:id/:username", handler.UserByUsername)
}
