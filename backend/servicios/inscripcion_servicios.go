package servicios

import (
	client "Arquitectura-de-Software-UCC/backend/clientes/inscripcion"
	"Arquitectura-de-Software-UCC/backend/dao"
	"Arquitectura-de-Software-UCC/backend/dominio"
	e "Arquitectura-de-Software-UCC/backend/utils"
)

type inscripcionServicio struct {
	inscripcionCliente client.InscripcionClienteInterface
}

type inscripcionServiceInterface interface {
	CrearInscripcion(newcurso dominio.InscripcionData) (dominio.InscripcionData, e.ApiError)
}

var (
	InscripcionServicio inscripcionServiceInterface
)

func initInscripcionService(inscripcionCliente client.InscripcionClienteInterface) inscripcionServiceInterface {
	service := new(inscripcionServicio)
	service.inscripcionCliente = inscripcionCliente
	return service
}

func init() {
	InscripcionServicio = initInscripcionService(client.InscripcionCliente)
}

func (s *inscripcionServicio) CrearInscripcion(newinscripcion dominio.InscripcionData) (dominio.InscripcionData, e.ApiError) {
	var inscripcion dao.Inscripcion

	inscripcion.UsuarioID = newinscripcion.UsuarioID
	inscripcion.CursoID = newinscripcion.CursoID

	inscripcion = s.inscripcionCliente.CrearInscripcion(inscripcion)

	newinscripcion.ID = inscripcion.ID

	return newinscripcion, nil
}
