package servicios

import(
	client "Arquitectura-de-Software-UCC/backend/clientes/cursos"
	"Arquitectura-de-Software-UCC/backend/dao"
	"Arquitectura-de-Software-UCC/backend/dominio/cursos"
	e "Arquitectura-de-Software-UCC/backend/utils"

)

type cursoServicio struct{
	cursoCliente client.CursoClienteInterface
}

type cursoServiceInterface interface{
	CrearCurso(newcurso dominio.CursoData) (dominio.CursoData, e.ApiError)
}

var(
	CursoServicio cursoServiceInterface
)

func initCursoService(cursoCliente client.CursoClienteInterface) cursoServiceInterface{
	service := new(cursoServicio)
	service.cursoCliente = cursoCliente
	return service
}

func init(){
	CursoServicio = initCursoService(client.CursoCliente)
}

func (s *cursoServicio) CrearCurso(newcurso dominio.CursoData) (dominio.CursoData, e.ApiError){
	var cu dao.Curso
	
	cu.Nombre = newcurso.Nombre
	cu.Descripcion = newcurso.Descripcion
	
	cu = s.cursoCliente.CrearCurso(cu)

	newcurso.CursoID = cu.CursoID

	return newcurso, nil
}