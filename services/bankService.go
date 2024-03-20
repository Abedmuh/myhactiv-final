package services

import (
	"myhactiv/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BankSvcInter interface {
	AddBank(Bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error)
	UpdateBank(Bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error)
	DeleteBank(ID uint, tx *gorm.DB, c *gin.Context) error

	AuthoBank(id string, tx *gorm.DB, c *gin.Context) error
}

type BankService struct {
}

func NewBankService() BankSvcInter {
	return &BankService{}
}

func (bs *BankService) AddBank(Bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error) {
  // tx.Create(&Bank)
  return models.Bank{}, nil
}

func (bs *BankService) UpdateBank(Bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error) {
  // tx.Save(&Bank)
  return Bank, nil
}

func (bs *BankService) DeleteBank(ID uint, tx *gorm.DB, c *gin.Context) error {
  // tx.Delete(&ID)
  return nil
}

func (bs *BankService) AuthoBank(id string, tx *gorm.DB, c *gin.Context) error {
  return nil
}