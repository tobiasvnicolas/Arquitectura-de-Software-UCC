package main

import (
	"backend/app"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	app.MapRoutes(engine) //mapea las rutas de la aplicación
	engine.Run(":8080")   //40:00 Explicación de quien llama a quien
	enginer := 4
}

//47:15 Explicación BindJson
//el Domain se usa para declarar las estructuras
///53:40 explicación de como pasarle el dato alusuario desde el controller
//jdnfugbrd
