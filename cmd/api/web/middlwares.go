package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		// "http://127.0.0.1:8844",
		// "http://localhost:8844",
		// "http://127.0.0.1:8877",
		// "http://localhost:8877",
		"*",
	}
	return cors.New(config)
}
