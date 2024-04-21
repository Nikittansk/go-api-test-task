package routes

import (
	"github.com/Nikittansk/go-api-test-task/config"
	"github.com/Nikittansk/go-api-test-task/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	api := router.Group(cfg.Endpoint)
	{
		version1 := api.Group("/v1")
		{
			version1.GET("/users", controllers.GetUsers)
			version1.GET("/users/:id", controllers.GetUser)
			version1.POST("/users", controllers.InsertUser)
			version1.PUT("/users/:id", controllers.UpdateUser)
			version1.DELETE("/users/:id", controllers.DeleteUser)

			version1.GET("/comments", controllers.GetComments)
			version1.GET("/comments/:id", controllers.GetComment)
			version1.PUT("/comments/:id", controllers.UpdateComment)
			version1.DELETE("/comments/:id", controllers.DeleteComment)

			version1.GET("/user/comments", controllers.GetComments)
			version1.GET("/user/:id/comments", controllers.GetUserComments)
			version1.POST("/user/:id/comments", controllers.InsertUserComment)
		}
	}

	return router
}
