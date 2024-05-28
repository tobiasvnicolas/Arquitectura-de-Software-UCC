package app

import(
	usuarioControlador 	"Arquitectura-de-Software-UCC/backend/controlador/usuario"
	log "github.com/sirupsen/logrus"

)

func mapUrls(){

	router.POST("/login", usuarioControlador.Login)
	router.GET("/usuario/:email", usuarioControlador.GetUsuariobyEmail)
	router.POST("/usuario", usuarioControlador.CrearUsuario)
	log.Info("Finishing mappings configurations")

}