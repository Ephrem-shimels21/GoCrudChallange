package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Handle500 middleware catches internal server errors and returns a consistent response
func Handle500() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{
					"error": "Internal server error. Please try again later.",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
