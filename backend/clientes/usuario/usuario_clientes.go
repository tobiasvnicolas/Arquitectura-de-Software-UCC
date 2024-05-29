package clientes

import (
	"Arquitectura-de-Software-UCC/backend/dao"
	
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type usuarioCliente struct{}

type UsuarioClienteInterface interface{
	GetUsuariobyEmail (email string) (dao.Usuario, error)
	CrearUsuario(user dao.Usuario) dao.Usuario
}

var(
	UsuarioCliente UsuarioClienteInterface
)

func init(){
	UsuarioCliente = &usuarioCliente{}
}

func (s *usuarioCliente) GetUsuariobyEmail( email string) (dao.Usuario, error) {
	var usuario dao.Usuario

	result := Db.Where("email = ?", email).First(&usuario)
	if result.Error != nil {
		return usuario, result.Error
	}
	log.Debug("Usuario: ", usuario)

	return usuario, nil
}

func (s *usuarioCliente) CrearUsuario (user dao.Usuario) dao.Usuario{

	result := Db.Create(&user)

	if result.Error != nil {
		log.Error ("")
	}

	log.Debug("Usuario Creado: ", user.UsuarioID)
	return user

}