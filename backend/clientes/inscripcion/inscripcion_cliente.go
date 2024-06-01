package inscripcion

import (
	dao "Arquitectura-de-Software-UCC/backend/dao"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type inscripcionCliente struct{}

type InscripcionClienteInterface interface {
	CrearInscripcion(curso dao.Inscripcion) dao.Inscripcion
	GetInscripcionById(id int64) (dao.Inscripcion, error)
}

var (
	InscripcionCliente InscripcionClienteInterface
)

func init() {
	InscripcionCliente = &inscripcionCliente{}
}

func (s *inscripcionCliente) CrearInscripcion(curso dao.Inscripcion) dao.Inscripcion {
	result := Db.Create(&inscripcion)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Curso Creado: ", curso.CursoID)
	return curso
}
