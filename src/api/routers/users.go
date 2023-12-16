package routers

import (
	"github.com/alireza-fa/golang-car-shop/api/handlers"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUserHandler(cfg)

	router.POST("/send-otp", h.SendOtp)
}
