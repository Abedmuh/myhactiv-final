package services

import (
	"myhactiv/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BankSvcInter interface {
	AddBank(bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error)
	UpdateBank(bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error)
	DeleteBank(ID uint, tx *gorm.DB, c *gin.Context) error

	AuthoBank(id string, tx *gorm.DB, c *gin.Context) error
}

type BankService struct {
}

func NewBankService() BankSvcInter {
	return &BankService{}
}

func (bs *BankService) AddBank(bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error) {
	result := tx.Create(&bank)
	if result.Error != nil {
		return models.Bank{}, result.Error
	}
  return bank, nil
}

func (bs *BankService) UpdateBank(bank models.Bank, tx *gorm.DB, c *gin.Context) (models.Bank, error) {
  result := tx.Save(&bank)
	if result.Error!= nil {
    return models.Bank{}, result.Error
  }
  return bank, nil
}

func (bs *BankService) DeleteBank(ID uint, tx *gorm.DB, c *gin.Context) error {
	var bank models.Bank
  result :=tx.Unscoped().Delete(&bank, ID)

	if result.Error != nil {
		return result.Error
	}
  return nil
}

func (bs *BankService) AuthoBank(id string, tx *gorm.DB, c *gin.Context) error {
  return nil
}