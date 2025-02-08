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

	// Blog Post routes
	blogRoutes := r.Group("/blog-post")
	{
		blogRoutes.POST("/", blogPostHandler.CreateBlogPost)
		blogRoutes.GET("/", blogPostHandler.GetBlogPosts)
		blogRoutes.GET("/:ID", blogPostHandler.GetBlogPost)
		blogRoutes.DELETE("/:ID", blogPostHandler.DeleteBlogPost)
		blogRoutes.PATCH("/:ID", blogPostHandler.UpdateBlogPost)
	}

	return r
}
