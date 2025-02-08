package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/DurgeshKr2242/blogassessment/domains"
	"github.com/DurgeshKr2242/blogassessment/mock"
	"github.com/gin-gonic/gin"
)

func TestBlogPostHandler_GetBlogPost(t *testing.T) {
	server := gin.New()
	fakeDomain := &mock.FakeService{}

	handler := NewBlogPostHandler(fakeDomain)

	route := "/blog-post/:ID"
	routeHttpMethod := http.MethodGet

	server.Handle(routeHttpMethod, route, handler.GetBlogPost)
	httpServer := httptest.NewServer(server)

	cases := map[string]struct {
		ID       string
		err      mock.ErrMock
		status   int
		response gin.H
	}{
		"When blog post is retrived successfully": {
			ID:     mock.MockID.String(),
			err:    mock.OK,
			status: http.StatusOK,
			response: gin.H{
				"blog": gin.H{
					"id":          &mock.MockID,
					"body":        "Some body for the blog",
					"description": "Some description for the blog",
					"title":       "Some title for blog",
					"created_at":  "2025-02-07T22:01:38.640214Z",
					"updated_at":  "2025-02-07T22:01:38.640214Z",
				},
			},
		},
		"When blog post is not found": {
			ID:     mock.MockID.String(),
			err:    mock.DBNotFoundError,
			status: http.StatusNotFound,
			response: gin.H{
				"message": domains.ErrorBlogPostNotFound.Error(),
			},
		},
		"When blog post get call fails due to unknown reason": {
			ID:     mock.MockID.String(),
			err:    mock.DBOperationError,
			status: http.StatusInternalServerError,
			response: gin.H{
				"message": domains.ErrorGetBlogPostFailed.Error(),
			},
		},
		"When Blog ID is invalid": {
			ID:     "invalid-id",
			err:    mock.OK,
			status: http.StatusBadRequest,
			response: gin.H{
				"message": []gin.H{
					{
						"ID": "must be a valid UUID",
					},
				},
			},
		},
	}

	gin.SetMode(gin.TestMode)
	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			if v.err != mock.OK {
				fakeDomain.Err = v.err
			} else {
				fakeDomain.Err = mock.OK
			}

			client := http.Client{}
			requestURL := httpServer.URL + fmt.Sprintf("/blog-post/%s", v.ID)
			req, err := http.NewRequest(routeHttpMethod, requestURL, nil)
			if err != nil {
				t.Error("unexpected error: ", err)
			}

			res, err := client.Do(req)
			if err != nil {
				t.Error("unexpected error: ", err)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("unexpected error: ", err)
			}

			if status := res.StatusCode; status != v.status {
				t.Errorf("handler returned wrong status code: \ngot %v\nwant %v\n", status, v.status)
			}

			if !reflect.DeepEqual(v.response, body) {
				if v.status == http.StatusOK {
					var got gin.H
					err = json.Unmarshal(body, &got)
					if err != nil {
						t.Fatal(err)
					}
					if fmt.Sprint(v.response) != fmt.Sprint(got) {
						t.Errorf("handler returned unexpected body: \ngot %v\nwant %v\n", got, v.response)
					}
				} else {
					var got gin.H
					err = json.Unmarshal(body, &got)
					if err != nil {
						t.Fatal(err)
					}
					if fmt.Sprint(v.response) != fmt.Sprint(got) {
						t.Errorf("handler returned unexpected body: \ngot %v\nwant %v\n", got, v.response)
					}
				}
			}

		})
	}

}

func TestBlogPostHandler_GetBlogPosts(t *testing.T) {
	server := gin.New()
	fakeDomain := &mock.FakeService{}

	handler := NewBlogPostHandler(fakeDomain)

	route := "/blog-post"
	routeHttpMethod := http.MethodGet

	server.Handle(routeHttpMethod, route, handler.GetBlogPosts)
	httpServer := httptest.NewServer(server)

	cases := map[string]struct {
		err      mock.ErrMock
		status   int
		response gin.H
	}{
		"When blog posts are retrived successfully": {
			err:    mock.OK,
			status: http.StatusOK,
			response: gin.H{
				"blogs": []gin.H{
					{
						"id":          &mock.MockID,
						"body":        "Some body for the blog",
						"description": "Some description for the blog",
						"title":       "Some title for blog",
						"created_at":  "2025-02-07T22:01:38.640214Z",
						"updated_at":  "2025-02-07T22:01:38.640214Z",
					},
				},
			},
		},
		"When blog posts get call fails due to unknown reason": {
			err:    mock.DBOperationError,
			status: http.StatusInternalServerError,
			response: gin.H{
				"message": domains.ErrorGetBlogPostsFailed.Error(),
			},
		},
	}

	gin.SetMode(gin.TestMode)
	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			if v.err != mock.OK {
				fakeDomain.Err = v.err
			} else {
				fakeDomain.Err = mock.OK
			}

			client := http.Client{}
			requestURL := httpServer.URL + "/blog-post"
			req, err := http.NewRequest(routeHttpMethod, requestURL, nil)
			if err != nil {
				t.Error("unexpected error: ", err)
			}

			res, err := client.Do(req)
			if err != nil {
				t.Error("unexpected error: ", err)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("unexpected error: ", err)
			}

			if status := res.StatusCode; status != v.status {
				t.Errorf("handler returned wrong status code: \ngot %v\nwant %v\n", status, v.status)
			}

			if !reflect.DeepEqual(v.response, body) {
				if v.status == http.StatusOK {
					var got gin.H
					err = json.Unmarshal(body, &got)
					if err != nil {
						t.Fatal(err)
					}
					if fmt.Sprint(v.response) != fmt.Sprint(got) {
						t.Errorf("handler returned unexpected body: \ngot %v\nwant %v\n", got, v.response)
					}
				} else {
					var got gin.H
					err = json.Unmarshal(body, &got)
					if err != nil {
						t.Fatal(err)
					}
					if fmt.Sprint(v.response) != fmt.Sprint(got) {
						t.Errorf("handler returned unexpected body: \ngot %v\nwant %v\n", got, v.response)
					}
				}
			}

		})
	}

}

// TestBlogPostHandler_CreateBlogPost tests the CreateBlogPost handler.
func TestBlogPostHandler_CreateBlogPost(t *testing.T) {
	server := gin.New()
	fakeDomain := &mock.FakeService{}

	handler := NewBlogPostHandler(fakeDomain)

	// We assume that update requests are made via PUT to the "/blog-post/:ID" route.
	route := "/blog-post"
	routeHttpMethod := http.MethodPost
	server.Handle(routeHttpMethod, route, handler.CreateBlogPost)
	httpServer := httptest.NewServer(server)

	cases := map[string]struct {
		body     gin.H
		Err      mock.ErrMock
		status   int
		response gin.H
	}{
		"When blog post is created successfully": {
			body: gin.H{
				"title":       "Created Title",
				"description": "Created description",
				"body":        "Created body",
			},
			Err:    mock.OK,
			status: http.StatusCreated,
			response: gin.H{
				"message": "blog post create successfully",
				"ID":      mock.MockID,
			},
		},
		"When create blog post call fails due to unknown reason": {
			body: gin.H{
				"title":       "Created Title",
				"description": "Created description",
				"body":        "Created body",
			},
			Err:    mock.DBOperationError,
			status: http.StatusInternalServerError,
			response: gin.H{
				"message": domains.ErrorCreateBlogPostFailed.Error(),
			},
		},
		"When req body is invalid": {
			body: gin.H{
				"title": "Cre",
				"body":  "Created body",
			},
			Err:    mock.DBOperationError,
			status: http.StatusBadRequest,
			response: gin.H{
				"message": []gin.H{
					{
						"Title": "should at least have 5 characters",
					},
					{
						"Description": "is required",
					},
				},
			},
		},
	}

	gin.SetMode(gin.TestMode)
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			fakeDomain.Err = tc.Err

			client := http.Client{}
			requestURL := fmt.Sprintf("%s/blog-post", httpServer.URL)

			jsonBody, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatal("failed to marshal JSON:", err)
			}

			req, err := http.NewRequest(routeHttpMethod, requestURL, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Error("unexpected error:", err)
			}
			req.Header.Set("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				t.Error("unexpected error:", err)
			}
			defer res.Body.Close()

			respBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("unexpected error reading body:", err)
			}

			if res.StatusCode != tc.status {
				t.Errorf("handler returned wrong status code:\ngot  %v\nwant %v\n", res.StatusCode, tc.status)
			}

			var got gin.H
			if err := json.Unmarshal(respBody, &got); err != nil {
				t.Fatal(err)
			}
			if fmt.Sprint(got) != fmt.Sprint(tc.response) {
				t.Errorf("handler returned unexpected body:\ngot  %v\nwant %v\n", got, tc.response)
			}
		})
	}
}

// TestBlogPostHandler_UpdateBlogPost tests the UpdateBlogPost handler.
func TestBlogPostHandler_UpdateBlogPost(t *testing.T) {
	server := gin.New()
	fakeDomain := &mock.FakeService{}

	handler := NewBlogPostHandler(fakeDomain)

	// We assume that update requests are made via PUT to the "/blog-post/:ID" route.
	route := "/blog-post/:ID"
	routeHttpMethod := http.MethodPatch
	server.Handle(routeHttpMethod, route, handler.UpdateBlogPost)
	httpServer := httptest.NewServer(server)

	cases := map[string]struct {
		id       string
		body     gin.H
		Err      mock.ErrMock
		status   int
		response gin.H
	}{
		"When blog post is updated successfully": {
			id: mock.MockID.String(),
			body: gin.H{
				"title":       "Updated Title",
				"description": "Updated description",
				"body":        "Updated body",
			},
			Err:    mock.OK,
			status: http.StatusOK,
			response: gin.H{
				"blog": gin.H{
					"id":          &mock.MockID,
					"title":       "Updated Title",
					"description": "Updated description",
					"body":        "Updated body",
					"created_at":  "2025-02-07T22:01:38.640214Z",
					"updated_at":  "2025-02-07T22:01:38.640214Z",
				},
			},
		},
		"When blog post is not found": {
			id: mock.MockID.String(),
			body: gin.H{
				"title":       "Updated Title",
				"description": "Updated description",
				"body":        "Updated body",
			},
			Err:    mock.DBNotFoundError,
			status: http.StatusNotFound,
			response: gin.H{
				"message": domains.ErrorBlogPostNotFound.Error(),
			},
		},
		"When update blog post call fails due to unknown reason": {
			id: mock.MockID.String(),
			body: gin.H{
				"title":       "Updated Title",
				"description": "Updated description",
				"body":        "Updated body",
			},
			Err:    mock.DBOperationErrorUpdateBlog,
			status: http.StatusInternalServerError,
			response: gin.H{
				"message": domains.ErrorUpdateBlogPostFailed.Error(),
			},
		},
		"When blog ID is invalid": {
			id: "invalidID",
			body: gin.H{
				"title":       "Updated Title",
				"description": "Updated description",
				"body":        "Updated body",
			},
			Err:    mock.OK,
			status: http.StatusBadRequest,
			response: gin.H{
				"message": []gin.H{
					{
						"ID": "must be a valid UUID",
					},
				},
			},
		},
		"When get blog post fails": {
			id: mock.MockID.String(),
			body: gin.H{
				"title":       "Updated Title",
				"description": "Updated description",
				"body":        "Updated body",
			},
			Err:    mock.DBOperationError,
			status: http.StatusInternalServerError,
			response: gin.H{
				"message": domains.ErrorGetBlogPostFailed.Error(),
			},
		},
		"When request body is invalid": {
			id: mock.MockID.String(),
			body: gin.H{
				"title": "Upd",
			},
			Err:    mock.OK,
			status: http.StatusBadRequest,
			response: gin.H{
				"message": []gin.H{
					{
						"Title": "should at least have 5 characters",
					},
				},
			},
		},
	}

	gin.SetMode(gin.TestMode)
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			fakeDomain.Err = tc.Err

			client := http.Client{}
			requestURL := fmt.Sprintf("%s/blog-post/%s", httpServer.URL, tc.id)

			jsonBody, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatal("failed to marshal JSON:", err)
			}

			req, err := http.NewRequest(routeHttpMethod, requestURL, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Error("unexpected error:", err)
			}
			req.Header.Set("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				t.Error("unexpected error:", err)
			}
			defer res.Body.Close()

			respBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("unexpected error reading body:", err)
			}

			if res.StatusCode != tc.status {
				t.Errorf("handler returned wrong status code:\ngot  %v\nwant %v\n", res.StatusCode, tc.status)
			}

			var got gin.H
			if err := json.Unmarshal(respBody, &got); err != nil {
				t.Fatal(err)
			}
			if fmt.Sprint(got) != fmt.Sprint(tc.response) {
				t.Errorf("handler returned unexpected body:\ngot  %v\nwant %v\n", got, tc.response)
			}
		})
	}
}

// TestBlogPostHandler_DeleteBlogPost tests the DeleteBlogPost handler.
func TestBlogPostHandler_DeleteBlogPost(t *testing.T) {
	server := gin.New()
	fakeDomain := &mock.FakeService{}

	handler := NewBlogPostHandler(fakeDomain)

	route := "/blog-post/:ID"
	routeHttpMethod := http.MethodDelete
	server.Handle(routeHttpMethod, route, handler.DeleteBlogPost)
	httpServer := httptest.NewServer(server)

	cases := map[string]struct {
		id       string
		Err      mock.ErrMock
		status   int
		response gin.H
	}{
		"When blog post is deleted successfully": {
			id:     mock.MockID.String(),
			Err:    mock.OK,
			status: http.StatusOK,
			response: gin.H{
				"message": "blog post deleted successfully",
			},
		},
		"When blog post is not found": {
			id:     mock.MockID.String(),
			Err:    mock.DBNotFoundError,
			status: http.StatusNotFound,
			response: gin.H{
				"message": domains.ErrorBlogPostNotFound.Error(),
			},
		},
		"When delete blog post call fails due to unknown reason": {
			id:     mock.MockID.String(),
			Err:    mock.DBOperationError,
			status: http.StatusInternalServerError,
			response: gin.H{
				"message": domains.ErrorDeleteBlogPostFailed.Error(),
			},
		},
		"When blog ID is invalid": {
			id:     "invalidID",
			Err:    mock.OK,
			status: http.StatusBadRequest,
			response: gin.H{
				"message": []gin.H{
					{
						"ID": "must be a valid UUID",
					},
				},
			},
		},
	}

	gin.SetMode(gin.TestMode)
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			fakeDomain.Err = tc.Err

			client := http.Client{}
			requestURL := fmt.Sprintf("%s/blog-post/%s", httpServer.URL, tc.id)
			req, err := http.NewRequest(routeHttpMethod, requestURL, nil)
			if err != nil {
				t.Error("unexpected error:", err)
			}

			res, err := client.Do(req)
			if err != nil {
				t.Error("unexpected error:", err)
			}
			defer res.Body.Close()

			respBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error("unexpected error reading body:", err)
			}

			if res.StatusCode != tc.status {
				t.Errorf("handler returned wrong status code:\ngot  %v\nwant %v\n", res.StatusCode, tc.status)
			}

			var got gin.H
			if err := json.Unmarshal(respBody, &got); err != nil {
				t.Fatal(err)
			}
			if fmt.Sprint(got) != fmt.Sprint(tc.response) {
				t.Errorf("handler returned unexpected body:\ngot  %v\nwant %v\n", got, tc.response)
			}
		})
	}
}
