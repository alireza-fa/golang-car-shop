package main

import (
	"github.com/alireza-fa/golang-car-shop/api"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/cache"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"log"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatalf("%s", err)
	}

	api.InitialServer(cfg)
}
