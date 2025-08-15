package handlers

import (
	"cimrique-workerpool/internal/models"
	"context"
	"errors"
	"os"
	"testing"
	"time"
)

type MainTest struct {
	//what an app needs to have
	client              mockClientInterface            //mock client, only PullItem command that returns smt random
	productRepo         mockProductRepositoryInterface //mock all funcs to return nil so processrequest doesn't throw
	merchantRepo        mockMerchantRepositoryInterface
	merchantProductRepo mockMerchantProductRepositoryInterface
	ctx                 context.Context
	service             mockServiceInterface
	handler             Handler              //no mock
}

// mock interfaces & structs
type mockService struct {
	ctx                 context.Context
	client              mockClientInterface
	productRepo         mockProductRepositoryInterface
	merchantRepo        mockMerchantRepositoryInterface
	merchantProductRepo mockMerchantProductRepositoryInterface
}
type mockServiceInterface interface {
	CreateWorkers(workerCount int, queueName string)
	BlockWorkers()
}

type mockClient struct {
	FaultyItemCounter int
	ValidItemCounter  int

	FaultyItemSent int
	ValidItemSent  int
}
type mockClientInterface interface {
	PullItem(queueName string) (string, error)
	CheckSuccess() bool
}

type mockProductRepository struct{}
type mockProductRepositoryInterface interface {
	ReadID(ID int) *models.Product
	ReadName(name string) *models.Product
	AddProduct(req models.Request) error
	UpdateProduct(req models.Request) error
}

type mockMerchantRepository struct{}
type mockMerchantRepositoryInterface interface {
	ReadName(name string) *models.Merchant
}

type mockMerchantProductRepository struct{}
type mockMerchantProductRepositoryInterface interface {
	Read(merchantID int, productID int) *models.MerchantProduct
	AddMerchantProduct(req models.Request, productId int, merchantId int) error
	UpdateMerchantProduct(req models.Request, productId int, merchantId int) error
}

var MTest MainTest

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	MTest.ctx = ctx
	MTest.setupApp()
	code := m.Run()
	os.Exit(code)
}

func (m *MainTest) setupApp() {
	//setup main test with mocks
	m.productRepo = &mockProductRepository{}
	m.merchantRepo = &mockMerchantRepository{}
	m.merchantProductRepo = &mockMerchantProductRepository{}
	m.client = &mockClient{FaultyItemCounter: 20, ValidItemCounter: 20, FaultyItemSent: 0, ValidItemSent: 0}
	m.service = mockService{
		ctx: m.ctx, 
		client: m.client, 
		productRepo: m.productRepo, 
		merchantRepo: m.merchantRepo, 
		merchantProductRepo: m.merchantProductRepo,
	}
	m.handler = NewHandler(m.service)
}

// mock methods implemented

// mock service -> create & block workers same, callUpdateWorkers doesn't process request
func (s mockService) CreateWorkers(workerCount int, queueName string) {
	for range workerCount {
		go s.CallUpdateWorker(s.ctx, queueName)
	}
}
func (s mockService) BlockWorkers() {
	<-s.ctx.Done()
}
func (s mockService) CallUpdateWorker(ctx context.Context, queueName string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			//pull from queue
			_, err := s.client.PullItem(queueName)
			if err != nil {
				//failed to pull, skip item
				continue
			} else {
			}
		}
	}
}

func (m *mockClient) PullItem(queueName string) (string, error) {

	if m.FaultyItemCounter > 0 {
		//return faulty item
		m.FaultyItemCounter -= 1
		m.FaultyItemSent += 1
		return "", errors.New("")
	} else if m.ValidItemCounter > 0 {
		//return valid item
		m.ValidItemCounter -= 1
		m.ValidItemSent += 1
		return "", nil
	} else {
		return "", nil
	}
}

func (m *mockClient) CheckSuccess() bool {
	if m.FaultyItemCounter == 0 && m.ValidItemCounter == 0 && m.FaultyItemSent == 20 && m.ValidItemSent == 20 {
		return true
	}
	return false
}

func (m *mockProductRepository) ReadID(ID int) *models.Product {
	return nil
}

func (m *mockProductRepository) ReadName(name string) *models.Product {
	return nil
}

func (m *mockProductRepository) AddProduct(req models.Request) error {
	return nil
}

func (m *mockProductRepository) UpdateProduct(req models.Request) error {
	return nil
}

func (m *mockMerchantRepository) ReadName(name string) *models.Merchant {
	return nil
}

func (m *mockMerchantProductRepository) Read(merchantID int, productID int) *models.MerchantProduct {
	return nil
}

func (m *mockMerchantProductRepository) AddMerchantProduct(req models.Request, productId int, merchantId int) error {
	return nil
}

func (m *mockMerchantProductRepository) UpdateMerchantProduct(req models.Request, productId int, merchantId int) error {
	return nil
}
