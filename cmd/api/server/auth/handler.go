package auth

import (
	"blog/pkg/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Pong",
		})
	}
}

func RegisterAuthor(authService AuthenticationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		author := domain.Author{}
		if err := c.Bind(&author); err != nil {
			fmt.Println(err)
		}
		if err := authService.Register(author); err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user_created": true,
		})
	}
}

func Login(authService AuthenticationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		credentials := make(map[string]string)

		if err := c.Bind(&credentials); err != nil {
			fmt.Println(err)
		}

		token, err := authService.Login(credentials["email"], credentials["password"], c.ClientIP())

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func TokenForVisiting(authService AuthenticationService) gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println(context.GetHeader("X-Forwarded-For"))

		context.JSON(http.StatusOK, gin.H{
			"token": authService.TokenForVisiting(context.ClientIP()),
		})

	}
}
