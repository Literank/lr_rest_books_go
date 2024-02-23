package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"literank.com/rest-books/model"
)

// Initialize a database instance
var db *gorm.DB

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Initialize database
	initDB()

	// Define a health endpoint handler
	r.GET("/", func(c *gin.Context) {
		// Return a simple response indicating the server is healthy
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBook)
	r.POST("/books", createBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	// Run the server on port 8080
	r.Run(":8080")
}

// Initialize database
func initDB() {
	// Open SQLite database
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate
	db.AutoMigrate(&model.Book{})
}

// Get all books
func getBooks(c *gin.Context) {
	var books []model.Book
	db.Find(&books)

	c.JSON(http.StatusOK, books)
}

// Get single book
func getBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// Create a new book
func createBook(c *gin.Context) {
	var reqBody model.Book
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&reqBody)

	c.JSON(http.StatusCreated, reqBody)
}

// Update an existing book
func updateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var reqBody model.Book
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Updates(&reqBody)

	c.JSON(http.StatusOK, book)
}

// Delete an existing book
func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusNoContent, nil)
}
