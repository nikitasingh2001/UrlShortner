package controller

import (
	"net/http"
	"urlshortner/internal/constant"
	"urlshortner/internal/logger"
	"urlshortner/internal/models"
	"urlshortner/internal/service"

	"github.com/gin-gonic/gin"
)

func ShortTheUrl(c *gin.Context) {
	var shortUrlBody models.ShortUrlBody
	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		logger.Log.Error("Error in binding: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": constant.BindError})
		return
	}

	resp, err := service.ShortenURL(shortUrlBody.LongUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "data": resp, "short_url": resp.ShortUrl})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	record, err := service.GetLongURL(code)
	if err != nil {
		logger.Log.Error("Error in getting the longURL ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)
}
