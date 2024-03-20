package services

import (
	"myhactiv/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductSvcInter interface {
	AddProduct(Product models.Product, tx *gorm.DB, c *gin.Context) (models.Product, error)
	UpdateProduct(Product models.Product, tx *gorm.DB, c *gin.Context) (models.Product, error)
	DeleteProduct(ID uint, tx *gorm.DB, c *gin.Context) error

	GetProduct(ID uint, tx *gorm.DB, c *gin.Context) (models.Product, error)
	GetProducts(tx *gorm.DB, c *gin.Context) ([]models.Product, error)

	AuthoProduct(id string, tx *gorm.DB, c *gin.Context) error
}

type ProductService struct {
}

func NewProductService() ProductSvcInter {
	return &ProductService{}
}

func (bs *ProductService) AddProduct(Product models.Product, tx *gorm.DB, c *gin.Context) (models.Product, error) {
  // tx.Create(&Product)
  return models.Product{}, nil
}

func (bs *ProductService) UpdateProduct(Product models.Product, tx *gorm.DB, c *gin.Context) (models.Product, error) {
  // tx.Save(&Product)
  return Product, nil
}

func (bs *ProductService) DeleteProduct(ID uint, tx *gorm.DB, c *gin.Context) error {
  // tx.Delete(&ID)
  return nil
}

func (bs *ProductService) AuthoProduct(id string, tx *gorm.DB, c *gin.Context) error {
  return nil
}

func (bs *ProductService) GetProduct(ID uint, tx *gorm.DB, c *gin.Context) (models.Product, error) {
	return models.Product{}, nil
}

func (bs *ProductService) GetProducts(tx *gorm.DB, c *gin.Context) ([]models.Product, error) {
  return []models.Product{}, nil
}