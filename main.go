package main

import (
	"fmt"

	"github.com/nandani2203/url_shortener/handler"

	"github.com/nandani2203/url_shortener/store"
	"github.com/nandani2203/url_shortener/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the URL Shortener service...")
	// Initialize the Redis Store
	store.InitializeStore()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the High-Scale URL Shortener API",
		})
	})

	// Route to create a short URL
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/test-shorten", func(c *gin.Context) {
		exampleUrl := "https://github.com/nandani2203"
		shortCode := utils.GenerateShortLink(exampleUrl)
		c.JSON(200, gin.H{
			"original_url": exampleUrl,
			"short_code":   shortCode,
		})
	})

	// Route to handle redirection
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	fmt.Println("Server is running on port 9808...")
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
