package web

import (
	"github.com/gin-gonic/gin"
)


func CORS() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Header(
			"Access-Control-Allow-Origin",
			"http://127.0.0.1:8844",
		)
		c.Next()
	}
}