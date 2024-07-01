package api

import (
	"context"
	"coupon_service/internal/config"
	"coupon_service/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type API struct {
	srv *http.Server
	mux *gin.Engine
	svc service.Service
}

func New(options ...func(*API)) *API {
	api := &API{}

	for _, o := range options {
		o(api)
	}
	return api
}

func WithDefaultGinRouter() func(*API) {
	return func(a *API) {
		gin.SetMode(gin.ReleaseMode)
		r := new(gin.Engine)
		r = gin.New()
		r.Use(gin.Recovery())
		a.mux = r
	}
}

func WithCustomMiddleware() func(*API){
	return func(a *API) {
		a.mux.Use(ErrorHandler())
	}
}

func WithServer(cfg *config.Config, svc service.Service) func(*API) {
	return func(a *API) {
		a.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
			Handler: a.mux,
		}
		a.svc = svc
	}
}

func WithRoutes() func(*API) {
	return func(a *API) {
		apiGroup := a.mux.Group("/api")
		apiGroup.POST("/apply", a.Apply)
		apiGroup.POST("/create", a.Create)
		apiGroup.GET("/coupons", a.Get)
	}
}

func (a API) Start() {
	if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func (a API) Close() {
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh
	log.Println("shutdown signal received...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	log.Println("shutdown completed")
}
