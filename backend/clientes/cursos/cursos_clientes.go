package clientes

import(
"Arquitectura-de-Software-UCC/backend/dao"
"github.com/jinzhu/gorm"
log "github.com/sirupsen/logrus"
)


var Db *gorm.DB

type cursoCliente struct{}

type CursoClienteInterface interface{
	CrearCurso(curso dao.Curso) dao.Curso
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