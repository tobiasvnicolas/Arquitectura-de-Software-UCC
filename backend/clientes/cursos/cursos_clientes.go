package clientes

import (
	dao "Arquitectura-de-Software-UCC/backend/dao"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type cursoCliente struct{}

type CursoClienteInterface interface {
	CrearCurso(curso dao.Curso) dao.Curso
	GetCursoById(id int64) (dao.Curso, error)
}

var (
	CursoCliente CursoClienteInterface
)

func init() {
	CursoCliente = &cursoCliente{}
}

func (s *cursoCliente) CrearCurso(curso dao.Curso) dao.Curso {
	result := Db.Create(&curso)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Curso Creado: ", curso.CursoID)
	return curso
}

func (s *cursoCliente) GetCursoById(id int64) (dao.Curso, error) {
	var curso dao.Curso

	result := Db.Where("curso_id = ?", id).First(&curso)

	if result.Error != nil {
		return curso, result.Error
	}

	log.Debug("Curso: ", curso)

	return curso, nil

}
