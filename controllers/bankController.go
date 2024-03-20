package controllers

import (
	"myhactiv/models"
	"myhactiv/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BankCtrlInter interface {
	PostBank(c *gin.Context)
	PatchBank(c *gin.Context) 
	DeleteBank(c *gin.Context)
}


type BankController struct {
	BankService services.BankSvcInter
	DB *gorm.DB
	Validate *validator.Validate
}

func NewBankController(BankService services.BankSvcInter, db *gorm.DB, validate *validator.Validate) BankCtrlInter {
  return &BankController{
		BankService: BankService,
		DB: db, 
		Validate: validate}
}

func (bc *BankController) PostBank(c *gin.Context) {
	var bank models.Bank

	userID,_ := c.Get("userid")
	id := uint(userID.(float64))
	bank.OwnerID = id


  if err := c.ShouldBindJSON(&bank); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := bc.Validate.Struct(bank); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	bank, err :=bc.BankService.AddBank(bank, bc.DB, c)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{
		"message": "Bank added",
	  "bank": bank,
  })
}

func (bc *BankController) PatchBank(c *gin.Context) {
	var bank models.Bank

	userID,_ := c.Get("userid")
	id := uint(userID.(float64))
	bank.OwnerID = id

  if err := c.ShouldBindJSON(&bank); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  if err := bc.Validate.Struct(bank); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
	bank, err :=bc.BankService.UpdateBank(bank, bc.DB, c)
  if err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{"message": "Bank updated"})
}

func (bc *BankController) DeleteBank(c *gin.Context) {
	bankID := c.Param("id")
  bankIDUint, err := strconv.Atoi(bankID)
	if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Bank ID"})
			return
	}
  if err := bc.BankService.DeleteBank(uint(bankIDUint), bc.DB, c); err!= nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }
  c.JSON(200, gin.H{"message": "Bank deleted"})
}