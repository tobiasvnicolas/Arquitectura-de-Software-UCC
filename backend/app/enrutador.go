package app

import

func MapRoutes(engine *gin.Engine) { //de parametro de las rutas
		engine.POST("/usuarios/login", users.Login)

}