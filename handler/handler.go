package handler

import (
	"net/http"

	"github.com/nandani2203/url_shortener/store"
	"github.com/nandani2203/url_shortener/utils"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := utils.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:9808/" + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	// If the URL isn't found in Redis
	if initialUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Oops! This short URL doesn't exist or has expired.",
		})
		return
	}
	c.Redirect(302, initialUrl)
}
