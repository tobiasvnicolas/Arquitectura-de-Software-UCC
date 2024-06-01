package app

import (
	cursoControlador "Arquitectura-de-Software-UCC/backend/controlador/cursos"
	inscripcionControlador "Arquitectura-de-Software-UCC/backend/controlador/inscripcion"
	usuarioControlador "Arquitectura-de-Software-UCC/backend/controlador/usuario"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	//USUARIOS
	router.POST("/login", usuarioControlador.Login)
	router.GET("/usuario/:email", usuarioControlador.GetUsuariobyEmail)
	router.GET("/usuario/:id", usuarioControlador.GetUsuariobyID)
	router.POST("/usuario", usuarioControlador.CrearUsuario)

	// CURSOS
	router.POST("/cursos", cursoControlador.CrearCurso)
	router.GET("/cursos/:id", cursoControlador.GetCursoById)

	// INSCRIPCION
	router.POST("/inscripcion", inscripcionControlador.CrearInscripcion)
	log.Info("Finishing mappings configurations")

}
