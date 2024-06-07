package servicios

import (
	client "Arquitectura-de-Software-UCC/backend/clientes/inscripcion"
	"Arquitectura-de-Software-UCC/backend/dao"
	"Arquitectura-de-Software-UCC/backend/dominio"
	e "Arquitectura-de-Software-UCC/backend/utils"
)

type inscripcionServicio struct {
	inscripcionCliente client.InscripcionClienteInterface
	cursoServicio cursoServiceInterface

}

type inscripcionServiceInterface interface {
	CrearInscripcion(newcurso dominio.InscripcionData) (dominio.InscripcionData, e.ApiError)
	GetInscripcionByUserId (id int64) ([]dominio.CursoData, e.ApiError)
}

var (
	InscripcionServicio inscripcionServiceInterface
)

func initInscripcionService(inscripcionCliente client.InscripcionClienteInterface, cursoServicio cursoServiceInterface) inscripcionServiceInterface {
	service := &inscripcionServicio{
        inscripcionCliente: inscripcionCliente,
        cursoServicio:      cursoServicio,
    }
	return service
}

func init() {
	InscripcionServicio = initInscripcionService(client.InscripcionCliente, CursoServicio)
}

func (s *inscripcionServicio) CrearInscripcion(newinscripcion dominio.InscripcionData) (dominio.InscripcionData, e.ApiError) {
	var inscripcion dao.Inscripcion

	inscripcion.UsuarioID = newinscripcion.UsuarioID
	inscripcion.CursoID = newinscripcion.CursoID

	inscripcion = s.inscripcionCliente.CrearInscripcion(inscripcion)

	newinscripcion.ID = inscripcion.ID

	return newinscripcion, nil
}

func (s *inscripcionServicio) GetInscripcionByUserId (id int64) ([]dominio.CursoData, e.ApiError){

	inscripciones, err := s.inscripcionCliente.GetInscripcionByUserId(id)

	if err != nil{
		return nil, e.NewBadRequestApiError("Inscripciones no encontradas")
	}

	var cursoids []int64
	for _, inscripcion := range inscripciones {
		cursoids = append (cursoids, inscripcion.CursoID)
	}

	cursos , err := s.cursoServicio.GetCursosByIds(cursoids)
	
	if err != nil{
		return nil, e.NewBadRequestApiError("Error")
	}
	return cursos, nil

}
