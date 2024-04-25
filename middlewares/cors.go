package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// CORS returns middleware for handling CORS
func CORS() gin.HandlerFunc {
	// Allow all origins with default CORS policy
	return cors.Default()
}