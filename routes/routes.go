package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/what-crud/controllers"
	"github.com/what-crud/middlewares"
)

func ApiRoutes() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.Use(middlewares.CorsConfig())
	r.MaxMultipartMemory = 8 << 20

	route := r.Group("/api")
	{
		// auth routes
		users := route.Group("/users")
		{
			users.GET("", RateLimit(time.Minute, 30), middlewares.AuthMiddlewareJWT(), controllers.GetUsers)
			users.GET("/:id", RateLimit(time.Minute, 30), middlewares.AuthMiddlewareJWT(), controllers.GetUserByID)
			users.POST("", RateLimit(time.Minute, 10), middlewares.AuthMiddlewareJWT(), controllers.StoreUser)
			users.PATCH("/:id", RateLimit(time.Minute, 10), middlewares.AuthMiddlewareJWT(), controllers.UpdateUser)
			users.DELETE("/:id", RateLimit(time.Minute, 20), middlewares.AuthMiddlewareJWT(), controllers.DestroyUser)
		}
		route.GET("/profile", RateLimit(time.Minute, 20), middlewares.AuthMiddlewareJWT(), controllers.GetUserProfile)

		// guest routes
		route.POST("/login", RateLimit(time.Minute, 10), controllers.Login)
		route.POST("/register", RateLimit(time.Minute, 10), controllers.Register)
	}

	return r
}
