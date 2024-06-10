package cursos

import(
"Arquitectura-de-Software-UCC/backend/dao"
"github.com/jinzhu/gorm"
log "github.com/sirupsen/logrus"
)


var Db *gorm.DB

type cursoCliente struct{}

type CursoClienteInterface interface{
	CrearCurso(curso dao.Curso) dao.Curso
	GetCursoById(id int64) (dao.Curso, error)
	SearchCursos(palabra string) ([]dao.Curso, error)
	GetCursosByIds(id []int64) ([]dao.Curso, error)
	GetCursos() ([]dao.Curso, error)

}

var(
	CursoCliente CursoClienteInterface
)

func init(){
	CursoCliente = &cursoCliente{}
}

func (s *cursoCliente) CrearCurso (curso dao.Curso) dao.Curso{
	result := Db.Create(&curso)

	if result.Error != nil{
		log.Error ("")
	}

	log.Debug("Curso Creado: ", curso.CursoID)
	return curso
}

func(s *cursoCliente) GetCursoById (id int64) (dao.Curso, error){
	var curso dao.Curso

	result := Db.Where ("curso_id = ?", id).First(&curso)

	if result.Error != nil {
		return curso, result.Error
	}

	log.Debug("Curso: ", curso)

	return curso, nil


}

 func (s *cursoCliente) SearchCursos (palabra string) ([]dao.Curso, error){

	var cursos []dao.Curso

	result := Db.Where ("nombre LIKE ? OR categoria LIKE ?", "%"+palabra+"%","%"+palabra+"%").Find(&cursos)

	if result.Error != nil{
		return cursos, result.Error
	}

	return cursos, nil

}


func (s *cursoCliente) GetCursosByIds(id []int64) ([]dao.Curso, error){

	var cursos []dao.Curso

	result := Db.Where ("curso_id IN (?)", id).Find(&cursos)

	if result.Error != nil{
		return cursos, result.Error
	}

	return cursos, nil

}


func (s *cursoCliente) GetCursos() ([]dao.Curso, error){

	var cursos []dao.Curso

	result := Db.Find(&cursos)

	if result.Error != nil{
		return cursos, result.Error
	}

	return cursos, nil

}

