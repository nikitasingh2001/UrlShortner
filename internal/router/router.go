package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"urlshortner/internal/config"
	"urlshortner/internal/controller"
	"urlshortner/internal/logger"
)

func ClientRoutes() *gin.Engine {

	r := gin.Default()
	r.Use(authMiddleware())
	r.POST("/v1/url/short", controller.ShortTheUrl)
	r.GET("/v1/url/url/:code", controller.RedirectURL)

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Failed to run server: %v", err)
	}

	return r
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		configPath := "conf/config.toml"
		authConfig, err := config.ReadAuthConfig(configPath)

		if err != nil {
			logger.Log.Error("Error reading auth config:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Authentication failed"})
			c.Abort()
			return
		}

		username, password, ok := c.Request.BasicAuth()
		if !ok || username != authConfig.Auth.Username || password != authConfig.Auth.Password {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Authentication failed, Incorrect password"})
			c.Abort()
			return
		}
	}
}
