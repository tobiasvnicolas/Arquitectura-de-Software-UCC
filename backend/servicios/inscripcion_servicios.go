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
	GetInscripcionById(id int64) (dominio.InscripcionData, e.ApiError)
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

	inscripcion.ID = newinscripcion.ID
	inscripcion.UsuarioID = newinscripcion.UsuarioID
	inscripcion.CursoID = newinscripcion.CursoID

	inscripcion = s.inscripcionCliente.CrearInscripcion(inscripcion)

	newinscripcion.ID = inscripcion.ID

	return newinscripcion, nil
}

func (s *inscripcionServicio) GetInscripcionById(id int64) (dominio.InscripcionData, e.ApiError) {
	var inscripcion dao.Inscripcion

	inscripcion, err := s.inscripcionCliente.GetInscripcionById(id)

	var ins dominio.InscripcionData

	if err != nil {
		return ins, e.NewBadRequestApiError("Inscripcion No Encontrada")
	}

	ins.ID = inscripcion.ID
	ins.UsuarioID = inscripcion.UsuarioID
	ins.CursoID = inscripcion.CursoID

	return ins, nil

}
