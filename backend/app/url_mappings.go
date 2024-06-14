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
	router.POST("/usuario", usuarioControlador.CrearUsuario)
	//router.GET("/usuario/miscursos", usuarioControlador.GetUsuarioMiscursos)

	// CURSOS
	router.POST("/cursos", cursoControlador.CrearCurso)
	router.GET("/cursos/:id", cursoControlador.GetCursoById)
	router.GET("/cursos/buscar", cursoControlador.SearchCursos)
	router.GET("/cursos/todos", cursoControlador.GetCursos)

	// INSCRIPCION
	router.POST("/inscripcion", inscripcionControlador.CrearInscripcion)
	router.GET("/inscripcion/:id", inscripcionControlador.GetInscripcionByUserId)
	log.Info("Finishing mappings configurations")

}
