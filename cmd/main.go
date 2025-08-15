package main

import (
	"cimrique-workerpool/internal/client"
	"cimrique-workerpool/internal/config"
	"cimrique-workerpool/internal/handlers"
	metric "cimrique-workerpool/internal/metrics"
	"cimrique-workerpool/internal/redis_client"
	"cimrique-workerpool/internal/repositories"
	"cimrique-workerpool/internal/service"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	//get configs
	AppConfig := config.LoadConfig()

	//get dependencies
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", AppConfig.DBParams.Host, AppConfig.DBParams.User, AppConfig.DBParams.Password, AppConfig.DBParams.Name, AppConfig.DBParams.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("couldn't connect to database", err)
	}
	defer close(db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	redisClient := redis_client.NewRedisClient(AppConfig.QueueParams.Address, AppConfig.QueueParams.Password, AppConfig.QueueParams.Number, AppConfig.QueueParams.Protocol)
	client := client.NewWorkerServiceClient(redisClient)
	productRepo := repositories.NewProductRepository(db)
	merchantRepo := repositories.NewMerchantRepository(db)
	merchantProductRepo := repositories.NewMerchantProductRepository(db)

	//metrics
	reg := prometheus.NewRegistry()
	metric := metric.NewMetric(reg)
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	service := service.NewServicesFuncs(ctx, client, productRepo, merchantRepo, merchantProductRepo, metric)
	handler := handlers.NewHandler(service)

	handler.HandleWorkers()

	// expose /metrics on the same Fiber app/port
	app.Get("/metrics", adaptor.HTTPHandler(promHandler))
}

func close(db *gorm.DB) {
	db_, err := db.DB()
	if err != nil {
		panic(err)
	}
	db_.Close()
}
