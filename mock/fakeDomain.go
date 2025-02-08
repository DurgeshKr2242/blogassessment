package mock

import (
	"github.com/DurgeshKr2242/blogassessment/domains"
	"github.com/DurgeshKr2242/blogassessment/models"
	"github.com/google/uuid"
)

type ErrMock int

const (
	// OK ...
	OK ErrMock = iota

	// DBOperationError ...
	DBOperationError

	// DBNotFoundError ...
	DBNotFoundError

	// DBOperationErrorUpdateBlog ...
	DBOperationErrorUpdateBlog
)

// FakeService is a fake struct for domain Service.
type FakeService struct {
	Err ErrMock
}

var (
	MockID       = uuid.MustParse("259c7e70-57b0-40d9-8fd1-20a7ed901fae")
	MockBlogPost = models.BlogPost{
		ID:          &MockID,
		Body:        "Some body for the blog",
		Description: "Some description for the blog",
		Title:       "Some title for blog",
		CreatedAt:   "2025-02-07T22:01:38.640214Z",
		UpdatedAt:   "2025-02-07T22:01:38.640214Z",
	}
	MockBlogPosts = []models.BlogPost{
		{
			ID:          &MockID,
			Body:        "Some body for the blog",
			Description: "Some description for the blog",
			Title:       "Some title for blog",
			CreatedAt:   "2025-02-07T22:01:38.640214Z",
			UpdatedAt:   "2025-02-07T22:01:38.640214Z",
		},
	}
)

func (s *FakeService) CreateBlogPost(blog *models.BlogPost) (*uuid.UUID, error) {
	if s.Err == DBOperationError {
		return nil, domains.ErrorCreateBlogPostFailed
	}

	return &MockID, nil
}

func (s *FakeService) GetBlogPost(ID *uuid.UUID) (*models.BlogPost, error) {
	if s.Err == DBOperationError {
		return nil, domains.ErrorGetBlogPostFailed
	}
	if s.Err == DBNotFoundError {
		return nil, domains.ErrorBlogPostNotFound
	}

	return &MockBlogPost, nil
}

func (s *FakeService) GetBlogPosts() ([]models.BlogPost, error) {
	if s.Err == DBOperationError {
		return nil, domains.ErrorGetBlogPostsFailed
	}

	return MockBlogPosts, nil
}

func (s *FakeService) UpdateBlogPost(post *models.BlogPost) error {
	if s.Err == DBOperationErrorUpdateBlog {
		return domains.ErrorUpdateBlogPostFailed
	}

	post.ID = &MockID
	post.CreatedAt = "2025-02-07T22:01:38.640214Z"
	post.UpdatedAt = "2025-02-07T22:01:38.640214Z"
	return nil
}

func (s *FakeService) DeleteBlogPost(ID *uuid.UUID) error {
	if s.Err == DBOperationError {
		return domains.ErrorDeleteBlogPostFailed
	}
	if s.Err == DBNotFoundError {
		return domains.ErrorBlogPostNotFound
	}
	return nil
}
