package api

import (
	"fmt"
	"github.com/alireza-fa/golang-car-shop/api/middlewares"
	"github.com/alireza-fa/golang-car-shop/api/routers"
	"github.com/alireza-fa/golang-car-shop/api/validations"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/docs"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
	"github.com/alireza-fa/golang-car-shop/pkg/metrics"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitialServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidator()
	RegisterPrometheus()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.Prometheus())
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler) /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())

	RegisterRouter(r)
	RegisterSwagger(r, cfg)

	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

func RegisterRouter(r *gin.Engine) {
	conf := config.GetConfig()
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// Test
		health := v1.Group("/health")
		test := v1.Group("/test")

		// Users
		users := v1.Group("/users")

		// Properties
		propertyCategories := v1.Group("/property-categories", middlewares.Authentication(conf), middlewares.Authorization([]string{"admin"}))
		properties := v1.Group("/properties", middlewares.Authentication(conf), middlewares.Authorization([]string{"admin"}))

		// Base
		countries := v1.Group("/countries", middlewares.Authentication(conf), middlewares.Authorization([]string{"admin"}))
		cities := v1.Group("/cities", middlewares.Authentication(conf), middlewares.Authorization([]string{"admin"}))
		files := v1.Group("/files", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		companies := v1.Group("/companies", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		colors := v1.Group("/colors", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		years := v1.Group("/years", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))

		// Car
		carTypes := v1.Group("/car-types", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		gearboxes := v1.Group("/gearboxes", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		carModels := v1.Group("/car-models", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		carModelColors := v1.Group("/car-model-colors", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		carModelYears := v1.Group("/car-model-years", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		carModelPriceHistories := v1.Group("/car-model-price-histories", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		carModelImages := v1.Group("/car-model-images", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		carModelProperties := v1.Group("/car-model-properties", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))
		carModelComments := v1.Group("/car-model-comments", middlewares.Authentication(conf), middlewares.Authorization([]string{"default"}))

		// Test
		routers.Health(health)
		routers.TestRouter(test)

		// User
		routers.User(users, conf)

		// Base
		routers.Country(countries, conf)
		routers.City(cities, conf)
		routers.File(files, conf)
		routers.Company(companies, conf)
		routers.Color(colors, conf)
		routers.Year(years, conf)

		// Property
		routers.PropertyCategory(propertyCategories, conf)
		routers.Property(properties, conf)

		// Car
		routers.CarType(carTypes, conf)
		routers.Gearbox(gearboxes, conf)
		routers.CarModel(carModels, conf)
		routers.CarModelColor(carModelColors, conf)
		routers.CarModelYear(carModelYears, conf)
		routers.CarModelPriceHistory(carModelPriceHistories, conf)
		routers.CarModelImage(carModelImages, conf)
		routers.CarModelProperty(carModelProperties, conf)
		routers.CarModelComment(carModelComments, conf)

		r.Static("/static", "./uploads")

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
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

func RegisterPrometheus() {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}
}
