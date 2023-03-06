package main

import "github.com/gin-gonic/gin"

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var books []Book

func main() {
	router := gin.Default()

	books = append(books, Book{ID: "1", Title: "Book One", Author: "John Doe", Description: "This is book one"})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: "Steve Smith", Description: "This is book two"})

	router.GET("/books", getBook)
	router.GET("/books/:id", getBooks)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}

func getBook(c *gin.Context) {
	c.JSON(200, gin.H{"data": books})
}

func getBooks(c *gin.Context) {
	id := c.Param("id")
	for _, item := range books {
		if item.ID == id {
			c.JSON(200, gin.H{"data": item})
			return
		}
	}
	c.JSON(404, gin.H{"data": "Not found"})
}

func createBook(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	books = append(books, book)

	c.JSON(201, gin.H{"data": book})
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for i, item := range books {
		if item.ID == id {
			books[i] = book
			c.JSON(200, gin.H{"data": book})
			return
		}
	}
	c.JSON(404, gin.H{"data": "Not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	c.JSON(204, nil)
}
