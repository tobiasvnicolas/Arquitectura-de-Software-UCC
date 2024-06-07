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
	GetInscripcionByUserId(id int64) ([]dao.Inscripcion, error)
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

func (s *inscripcionCliente) GetInscripcionByUserId (id int64) ([]dao.Inscripcion, error){

	var inscripciones []dao.Inscripcion

	result := Db.Where("usuario_id = ?", id).Find(&inscripciones)

	if result.Error != nil{
		return inscripciones, result.Error
	}

	return inscripciones, nil

}
