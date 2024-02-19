package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string  `json:"id"`
	Genre  string  `json:"genre"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Image  string  `json:"image"`
	Price  float64 `json:"price"`
}

var books []Book

func main() {
	if err := retrieveData(); err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/book", addBook)
	router.PUT("/books/:id", updateBookByID)
	router.DELETE("/book/:id", removeBookByID)
	router.Run("localhost:5000")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	router.Use(cors.New(config))
}

func retrieveData() error {
	data, err := ioutil.ReadFile("books.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &books); err != nil {
		return err
	}
	return nil
}

func getBooks(c *gin.Context) {
	response := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Books   []Book `json:"books"`
	}{
		Status:  "success",
		Message: "List of books",
		Books:   books,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this book could not be found"})
}

func addBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBookByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			if err := c.BindJSON(&books[i]); err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, books[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this book could not be updated"})
}

func removeBookByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book successfully removed"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "this book could not be removed"})
}
