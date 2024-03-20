package controllers

import (
	"myhactiv/models"
	"myhactiv/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductCtrlInter interface {
	PostProduct(c *gin.Context)
	PatchProduct(c *gin.Context) 
	DeleteProduct(c *gin.Context)

	GetProduct(c *gin.Context)
	GetProducts(c *gin.Context)
}


type ProductController struct {
	ProductService services.ProductSvcInter
	DB *gorm.DB
	Validate *validator.Validate
}

func NewProductController(ProductService services.ProductSvcInter, db *gorm.DB, validate *validator.Validate) ProductCtrlInter {
  return &ProductController{
		ProductService: ProductService,
		DB: db, 
		Validate: validate}
}

func (bc *ProductController) PostProduct(c *gin.Context) {
	var product models.Product

	userID,_ := c.Get("userid")
	id := uint(userID.(float64))
	product.OwnerID = id


  if err := c.ShouldBindJSON(&product); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := bc.Validate.Struct(product); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	product, err := bc.ProductService.AddProduct(product, bc.DB, c)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
		"message": "Product added",
	  "Product": product,
  })
}

func (bc *ProductController) PatchProduct(c *gin.Context) {
	var product models.Product

	userID,_ := c.Get("userid")
	id := uint(userID.(float64))
	product.OwnerID = id

  if err := c.ShouldBindJSON(&product); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := bc.Validate.Struct(product); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	product, err :=bc.ProductService.UpdateProduct(product, bc.DB, c)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{"message": "Product updated"})
}

func (bc *ProductController) DeleteProduct(c *gin.Context) {
	var Product models.Product
  if err := c.ShouldBindJSON(&Product); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := bc.Validate.Struct(Product); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := bc.ProductService.DeleteProduct(Product.ID, bc.DB, c); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{"message": "Product deleted"})
}

func (bc *ProductController) GetProduct(c *gin.Context) {
	productId := c.Param("id")
  productIdUint, err := strconv.Atoi(productId)
	if err != nil {
			c.JSON(400, gin.H{"error": "Invalid product ID"})
			return
	}
  product, err := bc.ProductService.GetProduct(uint(productIdUint), bc.DB, c)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
    "message": "Product found",
    "Product": product,
  })
}

func (bc *ProductController) GetProducts(c *gin.Context) {
	
  products, err := bc.ProductService.GetProducts( bc.DB, c)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
    "message": "Products found",
    "Products": products,
  })
}