package handler

import (
	"github.com/gin-gonic/gin"
	"hoang.com/url-shortener/shortener"
	"hoang.com/url-shortener/store"
	"net/http"
)

type UrlCreateRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context, service *store.StorageService) {
	var urlCreateRequest UrlCreateRequest
	if err := c.ShouldBindJSON(&urlCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl := shortener.GenerateShortLink(urlCreateRequest.LongUrl)
	service.SaveUrlMapping(c, shortUrl, urlCreateRequest.LongUrl)
	c.JSON(http.StatusOK, gin.H{
		"short_url": shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context, service *store.StorageService) {
	shortUrl := c.Query("url")
	initialUrl := service.RetrieveInitialUrl(c, shortUrl)
	c.Redirect(302, initialUrl)
}
