package clientes

import (
	"Arquitectura-de-Software-UCC/backend/dao"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUsuariobyEmail(string email) dao.Usuario {
	var usuario dao.Usuario

	Db.Where("email = ?", email).First(&usuario)
	log.Debug("Usuario: ", usuario)

	return usuario
}