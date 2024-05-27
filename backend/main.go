package main

import (
	"Arquitectura-de-Software-UCC/backend/app"
	"Arquitectura-de-Software-UCC/backend/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()

	//engine := gin.New()
	//enrutador.MapUrls(engine)
	//engine.Run(":8080")
}
