package repositories

import (
	"cimrique-workerpool/internal/models"

	"gorm.io/gorm"
)

// struct
type MerchantProductRepository struct {
	db *gorm.DB
}

// constructor
func NewMerchantProductRepository(db *gorm.DB) MerchantProductRepository {
	repo := new(MerchantProductRepository)
	repo.db = db
	return *repo
}

// read command
func (r MerchantProductRepository) Read(merchantID int, productID int) *models.MerchantProduct {
	res := &models.MerchantProduct{}
	err := r.db.First(res, "merchant_id = ? AND product_id = ?", merchantID, productID)
	if err.Error != nil {
		return nil
	}
	return res
}

func (r MerchantProductRepository) AddMerchantProduct(req models.Request, productId int, merchantId int) error {
	//create new merchantproduct from request
	merchantProduct := &models.MerchantProduct{
		ProductID: productId,
		MerchantID: merchantId, 
		MerchantPrice: *req.Price,
		MerchantStock: *req.Stock,
	}

	result := r.db.Create(merchantProduct)

	return result.Error
}

func (r MerchantProductRepository) UpdateMerchantProduct(req models.Request, productId int, merchantId int) error {
	res := r.db.Model(&models.MerchantProduct{}).Where("product_id = ? AND merchant_id = ?", productId, merchantId)
	res.Updates(models.MerchantProduct{MerchantPrice: *req.Price, MerchantStock: *req.Stock})
	return res.Error
}
