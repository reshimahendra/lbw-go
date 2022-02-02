package middleware

import "github.com/gin-gonic/gin"

// Security middleware
func Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		// X-XSS-Protection
		c.Writer.Header().Add("X-XSS-Protection", "1; mode=block")

		// HTTP Strict Transport Security
		c.Writer.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		// X-Frame-Options
		c.Writer.Header().Add("X-Frame-Options", "SAMEORIGIN")

		// X-Content-Type-Options
		c.Writer.Header().Add("X-Content-Type-Options", "nosniff")

		// Content Security Policy
		c.Writer.Header().Add("Content-Security-Policy", "default-src 'self';")

		// X-Permitted-Cross-Domain-Policies
		c.Writer.Header().Add("X-Permitted-Cross-Domain-Policies", "none")

		// Referrer-Policy
		c.Writer.Header().Add("Referrer-Policy", "no-referrer")

		// Feature-Policy
		c.Writer.Header().Add("Feature-Policy", "microphone 'none'; camera 'none'")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
