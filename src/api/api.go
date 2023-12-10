package api

import (
	"fmt"
	"github.com/alireza-fa/golang-car-shop/api/routers"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/gin-gonic/gin"
)

func InitialServer() {
	cnf := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		// v1.GET("/health", func(c *gin.Context){
		// 	c.JSON(http.StatusOK, "working!!")
		// 	return
		// })
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%d", cnf.Server.Port))
}
