package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/what-crud/controllers"
)

func NewApiRoutes() (r *gin.Engine) {
	r = gin.Default()

	route := r.Group("/api")
	{
		users := route.Group("/users")
		{
			users.GET("", controllers.GetUsers)
			users.GET("/:id", controllers.GetUserByID)
			users.POST("", controllers.StoreUser)
			users.PATCH("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DestroyUser)
		}

		// auth ?
	}

	return
}
