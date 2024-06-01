package inscripcion

import (
	"Arquitectura-de-Software-UCC/backend/dao"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type inscripcionCliente struct{}

type inscripcionClienteInterface interface {
	CrearInscripcion(inscripcion dao.Inscripcion) dao.Inscripcion
	GetInscripcionById(id int64) (dao.Inscripcion, error)
}

var (
	InscripcionCliente inscripcionCliente
)

func init() {
	inscripcionCliente = &inscripcionCliente{}
}
