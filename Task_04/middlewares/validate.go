package middlewares

import (
	"Task_04/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 验证token
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//1.从请求头中获取token
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "missing token"})
			return
		}
		//2.去掉"Bear "前缀
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		//解析并验证token
		var claims jwt.MapClaims = map[string]interface{}{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.SecrectKey), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "invalid token"})
					return
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "token expired"})
					return
				} else {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "invalid token"})
				}
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "internal error"})
			}

			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "invalid token"})
			return
		}

		//4.将解析出claims保存到上下文
		c.Set("claims", claims)
		c.Next()
	}
}
