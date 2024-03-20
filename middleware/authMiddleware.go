package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
				c.AbortWithStatusJSON(401, gin.H{"error": "Unauthenticated"})
				return
		}

		// Periksa apakah token memiliki prefix "Bearer "
		if !strings.HasPrefix(tokenString, "Bearer ") {
				c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token format"})
				return
		}

		// Ambil token setelah prefix "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		secretKey := viper.GetString("JWT_SECRET_KEY")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
		})
		if err != nil {
				fmt.Println(err)
				c.AbortWithStatusJSON(401, gin.H{"error": "Unauthenticated"})
				return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Ambil nilai userid dari token claims
				userid, ok := claims["userid"].(string)
				if !ok {
						c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token claims"})
						return
				}
				// Set nilai userid dalam konteks Gin
				c.Set("userid", userid)
		} else {
				c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
				return
		}

		c.Next()
	}
}

