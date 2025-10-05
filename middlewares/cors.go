package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/what-crud/initializers"
)

func CorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:           []string{initializers.AppURL},
		AllowMethods:           []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"Origin"},
		AllowBrowserExtensions: true,
		ExposeHeaders:          []string{"Content-Length"},
		AllowCredentials:       true,
		MaxAge:                 36 * time.Hour,
	})
}
