package app

import(
	usuarioControlador 	"Arquitectura-de-Software-UCC/backend/controlador/usuario"
	inscripcionControlador "Arquitectura-de-Software-UCC/backend/controlador/inscripcion"
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
	router.GET("/cursos/buscar/:palabra", cursoControlador.SearchCursos)


	// INSCRIPCION
	router.POST("/inscripcion", inscripcionControlador.CrearInscripcion)
	router.GET("/inscripcion/:id", inscripcionControlador.GetInscripcionByUserId)
	log.Info("Finishing mappings configurations")

}