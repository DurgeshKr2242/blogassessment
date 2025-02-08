package router

import (
	"github.com/DurgeshKr2242/blogassessment/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(blogPostHandler *handlers.BlogPostHandler) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Game routes
	gameRoutes := r.Group("/blog-post")
	{
		gameRoutes.POST("/", blogPostHandler.CreateBlogPost)
		gameRoutes.GET("/", blogPostHandler.GetBlogPosts)
		gameRoutes.GET("/:ID", blogPostHandler.GetBlogPost)
		gameRoutes.DELETE("/:ID", blogPostHandler.DeleteBlogPost)
		gameRoutes.PATCH("/:ID", blogPostHandler.UpdateBlogPost)
	}

	return r
}
