package middleware

import (
	"blog/pkg/domain"
	"errors"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()
	ae := &domain.APIError{}

	for _, err := range c.Errors {
		if errors.As(err, &ae) {
			c.JSON(ae.Code, gin.H{
				"message":       ae.Message,
				"status_detail": ae.Status,
			})
		}
	}
}
