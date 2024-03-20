package services

import (
	"errors"
	"myhactiv/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSvcInter interface {
	AddUser(user models.User, tx *gorm.DB, c *gin.Context) (models.ResUser, error)
	DeleteUser(username string, tx *gorm.DB, c *gin.Context) error
	LoginUser(user models.UserLogin, tx *gorm.DB, c *gin.Context) (models.ResUserLog, error)

	CheckUser(id string, tx *gorm.DB, c *gin.Context) error
}

type UserService struct {
}

func NewUserService() UserSvcInter {
	return &UserService{}
}

func (us *UserService) AddUser(user models.User, tx *gorm.DB, c *gin.Context) (models.ResUser, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
			return models.ResUser{}, err
	}
	user.Password = string(hashedPassword)

  if err := tx.Create(&user).Error; err != nil {
		return models.ResUser{}, err
	}
	res := models.ResUser{
    Username: user.Username,
		Name: 	user.Name,
    Email: user.Email,
		Role: user.Role,
	}
	return res, nil
}

func (us *UserService) DeleteUser(id string, tx *gorm.DB, c *gin.Context) error {
  if err := tx.Unscoped().Where("username = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) CheckUser(id string, tx *gorm.DB, c *gin.Context)  error {
	var user models.User
	if err := tx.Where("username = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return errors.New("user already exists")
}

func (us *UserService) LoginUser(user models.UserLogin, db *gorm.DB, c *gin.Context) (models.ResUserLog, error) {
	var dbUser models.User

	if err := db.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.ResUserLog{}, errors.New("pengguna tidak ditemukan")
		}
		return models.ResUserLog{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return models.ResUserLog{}, errors.New("password salah")
	}

	signedToken,err:= generateToken(dbUser.ID)
	if err!= nil {
    return models.ResUserLog{}, err
  }

	userLogin := models.ResUserLog{
		Username:    dbUser.Username,
		Name:    		 dbUser.Name,
		AccessToken: signedToken,
	}

	// Kembalikan pengguna jika login berhasil
	return userLogin, nil
}

func generateToken(id uint)  (string,error){
	secretKey := viper.GetString("JWT_SECRET_KEY")
	timeExp := viper.GetDuration("JWT_TIME_EXP")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": id,
    "exp": time.Now().Add(time.Duration(timeExp) * time.Hour).Unix(), // Token akan kedaluwarsa dalam 12 jam
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken,nil
}