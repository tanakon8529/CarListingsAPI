package routes

import (
	"daveslist-emdpcv/api/controllers"
	"daveslist-emdpcv/api/middlewares"
	"daveslist-emdpcv/api/settings"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	config := settings.LoadEnv()
	api := r.Group(config.ApiPath)
	{
		api.POST("/v1/auth", controllers.TokenEndpoint)
		api.GET("/v1/health", controllers.HealthCheck)

		// User routes
		api.POST("/v1/users", controllers.CreateUser)
		api.GET("/v1/users/:id", controllers.GetUser)
		api.PUT("/v1/users/:id", controllers.UpdateUser)
		api.DELETE("/v1/users/:id", controllers.DeleteUser)

		// Category routes
		api.POST("/v1/categories", controllers.CreateCategory)
		api.GET("/v1/categories/:id", controllers.GetCategory)
		api.PUT("/v1/categories/:id", controllers.UpdateCategory)
		api.DELETE("/v1/categories/:id", controllers.DeleteCategory)
		api.GET("/v1/categories", controllers.ListCategories)

		// Group routes that require authentication
		auth := api.Group("/")
		auth.Use(middlewares.TokenAuthMiddleware())
		{
			// Listing routes
			auth.GET("/v1/listings", controllers.ListListings)
			auth.POST("/v1/listings", controllers.CreateListing)
			auth.GET("/v1/listings/:id", controllers.GetListing)
			auth.PUT("/v1/listings/:id", controllers.UpdateListing)
			auth.DELETE("/v1/listings/:id", controllers.DeleteListing)

			// Reply routes
			auth.POST("/v1/listings/:listing_id/replies", controllers.CreateReply)
			auth.GET("/v1/replies/:id", controllers.GetReply)
			auth.PUT("/v1/replies/:id", controllers.UpdateReply)
			auth.DELETE("/v1/replies/:id", controllers.DeleteReply)

			// Private message routes
			auth.POST("/v1/messages", controllers.SendPrivateMessage)
			auth.GET("/v1/messages/:id", controllers.GetPrivateMessage)
			auth.DELETE("/v1/messages/:id", controllers.DeletePrivateMessage)
		}
	}
}
