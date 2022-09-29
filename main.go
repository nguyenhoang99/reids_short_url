package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hoang.com/url-shortener/handler"
	"hoang.com/url-shortener/store"
)

func main() {

	storageService := store.InitializeStore()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey go URL Shortener",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c, storageService)
	})

	r.GET("/short-url", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c, storageService)
	})

	err := r.Run(":9080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error : %v", err))
	}
}
