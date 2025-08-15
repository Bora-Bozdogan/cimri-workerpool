package service

import (
	"cimrique-workerpool/internal/models"
	"context"
	"encoding/json"
)

func (s ServicesFuncs) CallUpdateWorker(ctx context.Context, queueName string) {
	for {
		select {
		case <-ctx.Done():
			s.DecrementActiveWorkerCount()
			return
		default:
			//pull from queue
			res, err := s.client.PullItem(queueName)
			if err != nil {
				//failed to pull, skip item
				continue
			} else {
				req := &models.Request{}
				json.Unmarshal([]byte(res), req)
				s.ProcessRequest(*req)
			}
		}
	}
}

func (s ServicesFuncs) ProcessRequest(m models.Request) {
	//check if product exists, if no add, if yes, update
	product := s.productRepo.ReadName(m.ProductName)
	var err error

	if product == nil {
		//product nonexistent, create all fields
		err = s.productRepo.AddProduct(m)

		if err != nil {
			//dead letter que logic here
		}

		product = s.productRepo.ReadName(m.ProductName) //update product so no null pointer dereference at merchantproduct
	} else {
		//product exists, don't update popularity and urgency, update rest
		err = s.productRepo.UpdateProduct(m)

		if err != nil {
			//dead letter que logic here
		}
	}

	//check if merchantProduct exists, if no add, if yes, update
	var merchantProduct *models.MerchantProduct = nil
	merchant := s.merchantRepo.ReadName(m.StoreName)
	if merchant != nil && product != nil {
		merchantProduct = s.merchantProductRepo.Read(merchant.ID, product.ID)
	}

	if merchantProduct == nil {
		//merchantProduct nonexistent, write new merchantproduct
		err = s.merchantProductRepo.AddMerchantProduct(m, product.ID, merchant.ID)

		if err != nil {
			//dead letter que logic here
		}
	} else {
		//merchantproduct exists, update fields
		err = s.merchantProductRepo.UpdateMerchantProduct(m, product.ID, merchant.ID)

		if err != nil {
			//dead letter que logic here
		}
	}
}
