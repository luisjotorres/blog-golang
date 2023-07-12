package middleware

import (
	"blog/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

// ToDO obtener esta valor de las variables de entorno
var jwtKey = []byte("my_secret_key")

func Authenticate(c *gin.Context) {
	tokenStr := c.GetHeader("authorization")
	claims := &utils.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err.Error(), token.Valid)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Youre UnAuthorized",
			"status":  "USER_UNAUTHORIZED",
		})
		return
	}

	if claims.UserType == "author" {
		c.Set("userId", claims.UserId)
		c.Set("isVisit", false)
	} else {
		// Con esto mapearemos la ip del user y a donde entro
		c.Set("isVisit", true)
		c.Set("ip", claims.IpAddress)
		c.Set("action", c.FullPath())
	}
}
