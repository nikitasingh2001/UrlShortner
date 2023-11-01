package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"urlshortner/internal/config"
	"urlshortner/internal/controller"
)

func ClientRoutes() *gin.Engine {

	r := gin.Default()
	r.Use(authMiddleware())
	r.POST("/v1/url/short", controller.ShortTheUrl)
	r.GET("/v1/url/url/:code", controller.RedirectURL)

	applicationConfig := config.GetConfig()
	if err := r.Run(":" + strconv.Itoa(applicationConfig.Application.Port)); err != nil {
		log.Printf("Failed to run server: %v", err)
	}

	return r
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authConfig := config.GetConfig()

		username, password, ok := c.Request.BasicAuth()
		if !ok || username != authConfig.AuthUser.Username || password != authConfig.AuthUser.Password {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Authentication failed, Incorrect password"})
			c.Abort()
			return
		}
	}
}
