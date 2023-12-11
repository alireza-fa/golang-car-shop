package main

import (
	"github.com/alireza-fa/golang-car-shop/api"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/cache"
)

func main() {
	cfg := config.GetConfig()

	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitialServer(cfg)
}
