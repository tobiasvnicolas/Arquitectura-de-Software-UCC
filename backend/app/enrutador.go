package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()

		// Configurar CORS
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowCredentials = true
		config.AddAllowHeaders("Authorization")
		router.Use(cors.New(config))
}

func StartRoute() {
	mapUrls()

	log.Info("Starting server")
	
	router.Run(":8080")
}



