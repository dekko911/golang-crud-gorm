package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/what-crud/controllers"
	"github.com/what-crud/middlewares"
)

func ApiRoutes() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.Use(middlewares.CorsConfig())

	route := r.Group("/api")
	{
		// auth routes
		users := route.Group("/users")
		{
			users.GET("", middlewares.AuthMiddlewareJWT(), controllers.GetUsers)
			users.GET("/:id", middlewares.AuthMiddlewareJWT(), controllers.GetUserByID)
			users.POST("", middlewares.AuthMiddlewareJWT(), controllers.StoreUser)
			users.PATCH("/:id", middlewares.AuthMiddlewareJWT(), controllers.UpdateUser)
			users.DELETE("/:id", middlewares.AuthMiddlewareJWT(), controllers.DestroyUser)
		}

		// guest routes
		route.POST("/login", controllers.Login)
		route.POST("/register", controllers.Register)
	}

	return r
}
