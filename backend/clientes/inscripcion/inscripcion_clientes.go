package inscripcion

import (
	"Arquitectura-de-Software-UCC/backend/dao"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type inscripcionCliente struct{}

type InscripcionClienteInterface interface {
	CrearInscripcion(inscripcion dao.Inscripcion) dao.Inscripcion
	GetInscripcionById(id int64) (dao.Inscripcion, error)
	//tal vez no sea necesaria, pero por si acaso
}

var (
	InscripcionCliente InscripcionClienteInterface
)

func init() {
	InscripcionCliente = &inscripcionCliente{}
}

func (s *inscripcionCliente) CrearInscripcion(inscripcion dao.Inscripcion) dao.Inscripcion {
	result := Db.Create(&inscripcion)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Inscripcion Creada: ", inscripcion.ID)
	return inscripcion
}

func (s *inscripcionCliente) GetInscripcionById(id int64) (dao.Inscripcion, error) {
	var inscripcion dao.Inscripcion

	result := Db.Where("inscripcion_id = ?", id).First(&inscripcion)

	if result.Error != nil {
		return inscripcion, result.Error
	}

	log.Debug("Inscripcion: ", inscripcion)

	return inscripcion, nil

}
