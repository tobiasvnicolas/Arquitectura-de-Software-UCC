package app

import (
	"Arquitectura-de-Software-UCC/backend/controlador/usuario"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartRoute() {
	engine := gin.New()
	MapRoutes(engine)
	log.Info("Starting server")
	router.Run(":8080")
}
func MapRoutes(engine *gin.Engine) {
	engine.POST("usuario/registrarse", controlador.Registrarse)

}
