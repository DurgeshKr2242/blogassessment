package models

import (
	"github.com/google/uuid"
)

// BlogPost represents a blog post.
type BlogPost struct {
	ID          *uuid.UUID `json:"id"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Body        string     `json:"body" binding:"required"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

type UpdateBlogPostRequest struct {
	Title       *string `json:"title,omitempty" binding:"omitempty,min=5,max=60"`
	Description *string `json:"description,omitempty" binding:"omitempty,min=10,max=300"`
	Body        *string `json:"body,omitempty" binding:"omitempty,min=10"`
}

type CreateBlogPostRequest struct {
	Title       string `json:"title" binding:"required,min=5,max=60"`
	Description string `json:"description" binding:"required,min=10,max=300"`
	Body        string `json:"body" binding:"required,min=10"`
}
