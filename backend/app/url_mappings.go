package app

import(
	usuarioControlador 	"Arquitectura-de-Software-UCC/backend/controlador/usuario"
	cursoControlador 	"Arquitectura-de-Software-UCC/backend/controlador/cursos"

	log "github.com/sirupsen/logrus"

)

func mapUrls(){

	//USUARIOS
	router.POST("/login", usuarioControlador.Login)
	router.GET("/usuario/:email", usuarioControlador.GetUsuariobyEmail)
	router.POST("/usuario", usuarioControlador.CrearUsuario)

	// CURSOS
	router.POST("/cursos", cursoControlador.CrearCurso)
	router.GET("/cursos/:id", cursoControlador.GetCursoById)


	// INSCRIPCION

	log.Info("Finishing mappings configurations")

}