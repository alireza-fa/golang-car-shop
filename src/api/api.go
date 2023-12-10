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

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%d", cnf.Server.Port))
}
