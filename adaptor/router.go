package adaptor

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"literank.com/rest-books/application"
	"literank.com/rest-books/application/dto"
	"literank.com/rest-books/application/executor"
	"literank.com/rest-books/domain/model"
)

const (
	fieldID     = "id"
	fieldOffset = "o"
	fieldQuery  = "q"
)

// RestHandler handles all restful requests
type RestHandler struct {
	bookOperator   *executor.BookOperator
	reviewOperator *executor.ReviewOperator
	userOperator   *executor.UserOperator
}

func MakeRouter(wireHelper *application.WireHelper) (*gin.Engine, error) {
	rest := &RestHandler{
		bookOperator:   executor.NewBookOperator(wireHelper.BookManager(), wireHelper.CacheHelper()),
		reviewOperator: executor.NewReviewOperator(wireHelper.ReviewManager()),
		userOperator:   executor.NewUserOperator(wireHelper.UserManager()),
	}
	// Create a new Gin router
	r := gin.Default()

	// Define a health endpoint handler
	r.GET("/", func(c *gin.Context) {
		// Return a simple response indicating the server is healthy
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/books", rest.getBooks)
	r.GET("/books/:id", rest.getBook)
	r.POST("/books", rest.createBook)
	r.PUT("/books/:id", rest.updateBook)
	r.DELETE("/books/:id", rest.deleteBook)
	r.GET("/books/:id/reviews", rest.getReviewsOfBook)
	r.GET("/reviews/:id", rest.getReview)
	r.POST("/reviews", rest.createReview)
	r.PUT("/reviews/:id", rest.updateReview)
	r.DELETE("/reviews/:id", rest.deleteReview)

	userGroup := r.Group("/users")
	userGroup.POST("", rest.userSignUp)
	userGroup.POST("/sign-in", rest.userSignIn)
	return r, nil
}

// Get all books
func (r *RestHandler) getBooks(c *gin.Context) {
	offset := 0
	offsetParam := c.Query(fieldOffset)
	if offsetParam != "" {
		value, err := strconv.Atoi(offsetParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
		offset = value
	}
	books, err := r.bookOperator.GetBooks(c, offset, c.Query(fieldQuery))
	if err != nil {
		fmt.Printf("Failed to get books: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Get single book
func (r *RestHandler) getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(fieldID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	book, err := r.bookOperator.GetBook(c, uint(id))
	if err != nil {
		fmt.Printf("Failed to get the book with %d: %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get the book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Create a new book
func (r *RestHandler) createBook(c *gin.Context) {
	var reqBody model.Book
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := r.bookOperator.CreateBook(c, &reqBody)
	if err != nil {
		fmt.Printf("Failed to create: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to create"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// Update an existing book
func (r *RestHandler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(fieldID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var reqBody model.Book
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := r.bookOperator.UpdateBook(c, uint(id), &reqBody)
	if err != nil {
		fmt.Printf("Failed to update: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to update"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Delete an existing book
func (r *RestHandler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(fieldID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := r.bookOperator.DeleteBook(c, uint(id)); err != nil {
		fmt.Printf("Failed to delete: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to delete"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Get all book reviews
func (r *RestHandler) getReviewsOfBook(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param(fieldID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}
	books, err := r.reviewOperator.GetReviewsOfBook(c, uint(bookID), c.Query(fieldQuery))
	if err != nil {
		fmt.Printf("Failed to get reviews of book %d: %v\n", bookID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Get single review
func (r *RestHandler) getReview(c *gin.Context) {
	id := c.Param(fieldID)
	review, err := r.reviewOperator.GetReview(c, id)
	if err != nil {
		fmt.Printf("Failed to get the review %s: %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get the review"})
		return
	}
	c.JSON(http.StatusOK, review)
}

// Create a new review
func (r *RestHandler) createReview(c *gin.Context) {
	var reviewBody dto.ReviewBody
	if err := c.ShouldBindJSON(&reviewBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := r.reviewOperator.CreateReview(c, &reviewBody)
	if err != nil {
		fmt.Printf("Failed to create: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to create the review"})
		return
	}
	c.JSON(http.StatusCreated, review)
}

// Update an existing review
func (r *RestHandler) updateReview(c *gin.Context) {
	id := c.Param(fieldID)

	var reqBody model.Review
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := r.reviewOperator.UpdateReview(c, id, &reqBody)
	if err != nil {
		fmt.Printf("Failed to update: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to update the review"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Delete an existing review
func (r *RestHandler) deleteReview(c *gin.Context) {
	id := c.Param(fieldID)

	if err := r.reviewOperator.DeleteReview(c, id); err != nil {
		fmt.Printf("Failed to delete: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to delete the review"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (r *RestHandler) userSignUp(c *gin.Context) {
	var ucBody dto.UserCredential
	if err := c.ShouldBindJSON(&ucBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := r.userOperator.CreateUser(c, &ucBody)
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to sign up"})
		return
	}
	c.JSON(http.StatusCreated, u)
}

func (r *RestHandler) userSignIn(c *gin.Context) {
	var m dto.UserCredential
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := r.userOperator.SignIn(c, m.Email, m.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}
