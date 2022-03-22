package middleware

import "github.com/gin-gonic/gin"

func Verifysign() gin.HandlerFunc {
	return func(c *gin.Context) {
		timestamp := c.Request.Header.Get("timestamp")
		sign := c.Request.Header.Get("sign")

	}
}
