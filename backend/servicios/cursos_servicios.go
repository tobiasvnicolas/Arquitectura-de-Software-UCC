package servicios

import(
	client "Arquitectura-de-Software-UCC/backend/clientes/cursos"
	"Arquitectura-de-Software-UCC/backend/dao"
	"Arquitectura-de-Software-UCC/backend/dominio"
	e "Arquitectura-de-Software-UCC/backend/utils"
	"strings"

)

type cursoServicio struct{
	cursoCliente client.CursoClienteInterface
}

type cursoServiceInterface interface{
	CrearCurso(newcurso dominio.CursoData) (dominio.CursoData, e.ApiError)
	GetCursoById (id int64) (dominio.CursoData, e.ApiError)
	SearchCursos (palabra string) ([]dominio.CursoData, e.ApiError)
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
	var curso dao.Curso
	
	curso.Nombre = newcurso.Nombre
	curso.Descripcion = newcurso.Descripcion
	curso.Categoria = newcurso.Categoria
	
	curso = s.cursoCliente.CrearCurso(curso)

	newcurso.CursoID = curso.CursoID

	return newcurso, nil
}

func (s *cursoServicio) GetCursoById(id int64) (dominio.CursoData, e.ApiError){
	var curso dao.Curso

	curso, err := s.cursoCliente.GetCursoById(id)
	
	var cu dominio.CursoData

	if err != nil {
		return cu, e.NewBadRequestApiError("Curso No Encontrado")
	}

	cu.CursoID = curso.CursoID
	cu.Nombre = curso.Nombre
	cu.Descripcion = curso.Descripcion
	cu.Categoria = curso.Categoria

	return cu, nil


}

func (s *cursoServicio) SearchCursos (palabra string) ([]dominio.CursoData, e.ApiError){

	palabras := strings.TrimSpace(palabra)

	cursos, err := s.cursoCliente.SearchCursos(palabras)

	if err != nil{
		return nil, e.NewBadRequestApiError("Curso No Encontrado")
	}

	results := make([]dominio.CursoData, 0)
	for _, curso := range cursos{
			results = append(results, dominio.CursoData{

				CursoID: curso.CursoID,
				Nombre: curso.Nombre,
				Descripcion: curso.Descripcion,
				Categoria: curso.Categoria,
		})
	}

	return results, nil


}