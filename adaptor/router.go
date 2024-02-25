package adaptor

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"literank.com/rest-books/application"
	"literank.com/rest-books/application/executor"
	"literank.com/rest-books/domain/model"
)

// RestHandler handles all restful requests
type RestHandler struct {
	bookOperator *executor.BookOperator
}

func MakeRouter(wireHelper *application.WireHelper) (*gin.Engine, error) {
	rest := &RestHandler{
		bookOperator: executor.NewBookOperator(wireHelper.BookManager()),
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
	return r, nil
}

// Get all books
func (r *RestHandler) getBooks(c *gin.Context) {
	books, err := r.bookOperator.GetBooks(c)
	if err != nil {
		fmt.Printf("Failed to get books: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Get single book
func (r *RestHandler) getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
		return
	}
	book, err := r.bookOperator.GetBook(c, uint(id))
	if err != nil {
		fmt.Printf("Failed to get the book with %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get the book"})
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
		fmt.Printf("Failed to create: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to create"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// Update an existing book
func (r *RestHandler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
		return
	}

	var reqBody model.Book
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := r.bookOperator.UpdateBook(c, uint(id), &reqBody)
	if err != nil {
		fmt.Printf("Failed to update: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to update"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Delete an existing book
func (r *RestHandler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
		return
	}

	if err := r.bookOperator.DeleteBook(c, uint(id)); err != nil {
		fmt.Printf("Failed to delete: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
