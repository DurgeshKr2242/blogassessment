package domains

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DurgeshKr2242/blogassessment/models"
	"github.com/google/uuid"
)

// BlogPostDomain defines the operations for blog posts.
type BlogPostDomain interface {
	CreateBlogPost(blog *models.BlogPost) (*uuid.UUID, error)
	GetBlogPost(ID *uuid.UUID) (*models.BlogPost, error)
	GetBlogPosts() ([]models.BlogPost, error)
	UpdateBlogPost(post *models.BlogPost) error
	DeleteBlogPost(ID *uuid.UUID) error
}

type blogPostDomain struct {
	db *sql.DB
}

// NewBlogPostDomain returns a new BlogPostDomain.
func NewBlogPostDomain(db *sql.DB) BlogPostDomain {
	return &blogPostDomain{db: db}
}

var (
	ErrorBlogPostNotFound     = errors.New("blog post not found")
	ErrorGetBlogPostFailed    = errors.New("failed to get blog post")
	ErrorGetBlogPostsFailed   = errors.New("failed to get blog posts")
	ErrorCreateBlogPostFailed = errors.New("failed to create blog post")
	ErrorUpdateBlogPostFailed = errors.New("failed to update blog post")
	ErrorDeleteBlogPostFailed = errors.New("failed to delete blog post")
)

func (d *blogPostDomain) CreateBlogPost(blog *models.BlogPost) (*uuid.UUID, error) {
	var ID *uuid.UUID
	query := `
       INSERT INTO blog_posts (title, description, body, created_at, updated_at)
       VALUES ($1, $2, $3, $4, $5)
       RETURNING id
    `
	now := time.Now()
	err := d.db.QueryRow(query, blog.Title, blog.Description, blog.Body, now, now).
		Scan(&ID)
	if err != nil {
		return nil, ErrorCreateBlogPostFailed
	}
	return ID, nil
}

func (d *blogPostDomain) GetBlogPost(ID *uuid.UUID) (*models.BlogPost, error) {
	query := `
       SELECT id, title, description, body, created_at, updated_at
       FROM blog_posts
       WHERE id = $1
    `

	var blog models.BlogPost
	err := d.db.QueryRow(query, ID).
		Scan(&blog.ID, &blog.Title, &blog.Description, &blog.Body, &blog.CreatedAt, &blog.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrorBlogPostNotFound
	} else if err != nil {
		return nil, ErrorGetBlogPostFailed
	}
	return &blog, nil
}

func (d *blogPostDomain) GetBlogPosts() ([]models.BlogPost, error) {
	query := `
       SELECT id, title, description, body, created_at, updated_at
       FROM blog_posts
       ORDER BY created_at DESC
    `
	rows, err := d.db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, ErrorGetBlogPostsFailed
	}
	defer rows.Close()

	blogs := []models.BlogPost{}
	for rows.Next() {
		var blog models.BlogPost
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.Body, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			fmt.Println(err.Error())
			return nil, ErrorGetBlogPostsFailed
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, ErrorGetBlogPostsFailed
	}
	return blogs, nil
}

func (d *blogPostDomain) UpdateBlogPost(blog *models.BlogPost) error {
	query := `
       UPDATE blog_posts
       SET title = $1, description = $2, body = $3, updated_at = $4
       WHERE id = $5
    `
	now := time.Now()
	_, err := d.db.Exec(query, blog.Title, blog.Description, blog.Body, now, blog.ID)
	if err != nil {
		return ErrorUpdateBlogPostFailed
	}
	return nil
}

func (d *blogPostDomain) DeleteBlogPost(ID *uuid.UUID) error {
	query := `DELETE FROM blog_posts WHERE id = $1`
	result, err := d.db.Exec(query, ID)
	if err != nil {
		return ErrorDeleteBlogPostFailed
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return ErrorDeleteBlogPostFailed
	}
	if rowsAffected == 0 {
		return ErrorBlogPostNotFound
	}
	return nil
}
