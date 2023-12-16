package api

import (
	"fmt"
	"github.com/alireza-fa/golang-car-shop/api/middlewares"
	"github.com/alireza-fa/golang-car-shop/api/routers"
	"github.com/alireza-fa/golang-car-shop/api/validations"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitialServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidator()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRouter(r)
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validations.PasswordValidator, true)
	}
}

func RegisterRouter(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test := v1.Group("/test")
		users := v1.Group("/users")

		routers.Health(health)
		routers.TestRouter(test)
		routers.User(users, config.GetConfig())
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
