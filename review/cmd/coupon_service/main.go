package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"log"
)

func main() {

	cfg := config.Load()


	repo := memdb.New()
	svc := service.New(repo)
	api := api.New(
		api.WithDefaultGinRouter(),
		api.WithCustomMiddleware(),
		api.WithServer(cfg,svc),
		api.WithRoutes(),
	)

	log.Println("Starting Coupon service server..")
	go api.Start()
	api.Close()
}
