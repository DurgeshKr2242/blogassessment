package handlers

import (
	"errors"
	"net/http"

	"github.com/DurgeshKr2242/blogassessment/domains"
	"github.com/DurgeshKr2242/blogassessment/helpers"
	"github.com/DurgeshKr2242/blogassessment/models"
	"github.com/DurgeshKr2242/blogassessment/validation"
	"github.com/gin-gonic/gin"
)

// BlogPostHandler handles blog post endpoints.
type BlogPostHandler struct {
	domain domains.BlogPostDomain
}

// NewBlogPostHandler creates a new BlogPostHandler.
func NewBlogPostHandler(domain domains.BlogPostDomain) *BlogPostHandler {
	return &BlogPostHandler{domain: domain}
}

func (h *BlogPostHandler) CreateBlogPost(c *gin.Context) {
	var req models.CreateBlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": validation.CustomValidationError(err),
		})
		return
	}

	blog := models.BlogPost{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
	}

	blogID, err := h.domain.CreateBlogPost(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "blog post create successfully",
		"ID":      blogID,
	})
}

func (h *BlogPostHandler) GetBlogPost(c *gin.Context) {
	request := struct {
		ID string `uri:"ID" binding:"required,uuid"`
	}{}

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": validation.CustomValidationError(err),
		})
		return
	}

	blogID := helpers.ParseUUID(request.ID)

	blog, err := h.domain.GetBlogPost(blogID)
	if err != nil {
		if errors.Is(domains.ErrorBlogPostNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blog": blog,
	})
}

func (h *BlogPostHandler) GetBlogPosts(c *gin.Context) {
	blogs, err := h.domain.GetBlogPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
	})
}

func (h *BlogPostHandler) UpdateBlogPost(c *gin.Context) {
	request := struct {
		ID string `uri:"ID" binding:"required,uuid"`
	}{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": validation.CustomValidationError(err),
		})
		return
	}
	blogID := helpers.ParseUUID(request.ID)

	blog, err := h.domain.GetBlogPost(blogID)
	if err != nil {
		if errors.Is(domains.ErrorBlogPostNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var req models.UpdateBlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validation.CustomValidationError(err)})
		return
	}

	if req.Title != nil {
		blog.Title = *req.Title
	}
	if req.Description != nil {
		blog.Description = *req.Description
	}
	if req.Body != nil {
		blog.Body = *req.Body
	}

	if err := h.domain.UpdateBlogPost(blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}

func (h *BlogPostHandler) DeleteBlogPost(c *gin.Context) {
	request := struct {
		ID string `uri:"ID" binding:"uuid,required"`
	}{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": validation.CustomValidationError(err),
		})
		return
	}
	blogID := helpers.ParseUUID(request.ID)

	if err := h.domain.DeleteBlogPost(blogID); err != nil {
		if errors.Is(domains.ErrorBlogPostNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "blog post deleted successfully"})
}
