package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://127.0.0.1:8844",
		"http://localhost:8844",
	}
	return cors.New(config)
}
