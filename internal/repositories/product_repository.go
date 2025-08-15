package repositories

import (
	"cimrique-workerpool/internal/models"
	"gorm.io/gorm"
)

// struct
type ProductRepository struct {
	db *gorm.DB
}

// constructor
func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(ProductRepository)
	repo.db = db
	return *repo
}

// read command
func (r ProductRepository) ReadID(ID int) *models.Product {
	res := &models.Product{}
	err := r.db.First(res, "id = ?", ID)
	if err.Error != nil {
		return nil
	}
	return res
}

func (r ProductRepository) ReadName(name string) *models.Product {
	res := &models.Product{}
	err := r.db.First(res, "name = ?", name)
	if err.Error != nil {
		return nil
	}
	return res
}

func (r ProductRepository) AddProduct(req models.Request) error {
	//create new product from request
	product := &models.Product{
		ProductName: req.ProductName,
		ProductDescription: req.ProductDescription,
		ProductImage: req.ProductImage,
		Popularity_score: req.PopularityScore,
		Urgency_score: req.UrgencyScore,
	}

	//add to the right table
	res := r.db.Create(product)

	return res.Error
}

func (r ProductRepository) UpdateProduct(req models.Request) error {
	res := r.db.Model(&models.Product{}).Where("name = ?", req.ProductName)
	res.Updates(models.Product{ProductName: req.ProductName, ProductDescription: req.ProductDescription, ProductImage: req.ProductImage})
	return res.Error
}