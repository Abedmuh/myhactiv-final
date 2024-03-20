package middleware

import (
	"fmt"

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
		// if !strings.HasPrefix(tokenString, "Bearer ") {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token format"})
		// 	return
		// }

		// tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		
		secretKey := viper.GetString("JWT_SECRET_KEY")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("auth token claims")
			userid, ok := claims["userid"]
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

