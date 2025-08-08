package main

import (
	"cimrique-workerpool/internal/client"
	"cimrique-workerpool/internal/config"
	"cimrique-workerpool/internal/handlers"
	"cimrique-workerpool/internal/repositories"
	"cimrique-workerpool/internal/service"
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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
	client := client.NewWorkerServiceClient(AppConfig.QueueParams.Address, AppConfig.QueueParams.Password, AppConfig.QueueParams.Number, AppConfig.QueueParams.Protocol)
	productRepo := repositories.NewProductRepository(db)
	merchantRepo := repositories.NewMerchantRepository(db)
	merchantProductRepo := repositories.NewMerchantProductRepository(db)

	service := service.NewServicesFuncs(ctx, client, productRepo, merchantRepo, merchantProductRepo)
	handler := handlers.NewHandler(service)

	handler.HandleWorkers()

}

func close(db *gorm.DB) {
	db_, err := db.DB()
	if err != nil {
		panic(err)
	}
	db_.Close()
}
