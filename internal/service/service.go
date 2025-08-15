package service

import (
	"cimrique-workerpool/internal/models"
	"context"
)

type ClientServiceInterface interface {
	PullItem(queueName string) (string, error)
}

type ProductRepositoryInterface interface {
	ReadID(ID int) *models.Product
	ReadName(name string) *models.Product
	AddProduct(req models.Request) error
	UpdateProduct(req models.Request) error
}

type MerchantRepositoryInterface interface {
	ReadName(name string) *models.Merchant
}

type MerchantProductRepositoryInterface interface {
	Read(merchantID int, productID int) *models.MerchantProduct
	AddMerchantProduct(req models.Request, productId int, merchantId int) error
	UpdateMerchantProduct(req models.Request, productId int, merchantId int) error
}

type metricsInterface interface {
	IncrementActiveWorkerCount()
	DecrementActiveWorkerCount()
}

type ServicesFuncs struct {
	ctx                 context.Context
	client              ClientServiceInterface
	productRepo         ProductRepositoryInterface
	merchantRepo        MerchantRepositoryInterface
	merchantProductRepo MerchantProductRepositoryInterface
	metrics metricsInterface
}

func NewServicesFuncs(context context.Context, client ClientServiceInterface, productRepo ProductRepositoryInterface, merchantRepo MerchantRepositoryInterface, merchantProductRepo MerchantProductRepositoryInterface, metrics metricsInterface) ServicesFuncs {
	return ServicesFuncs{ctx: context, client: client, productRepo: productRepo, merchantRepo: merchantRepo, merchantProductRepo: merchantProductRepo, metrics: metrics}
}
