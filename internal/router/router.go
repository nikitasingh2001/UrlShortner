package router

import (
	"log"
	"os"

	"urlshortner/internal/controller"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func ClientRoutes() {

	r := routes{
		router: gin.Default(),
	}

	r.router.POST("/v1/url/short", controller.ShortTheUrl)
	r.router.GET("/v1/url//url/:code", controller.RedirectURL)

	if err := r.router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
