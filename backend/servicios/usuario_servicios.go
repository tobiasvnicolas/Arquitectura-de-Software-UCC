package servicios

import(
	client "Arquitectura-de-Software-UCC/backend/clientes/usuario"
	"Arquitectura-de-Software-UCC/backend/dao"
	"Arquitectura-de-Software-UCC/backend/dominio"
	"fmt"
	"errors"
	"strings"
	"crypto/md5"
	log "github.com/sirupsen/logrus"

	e "Arquitectura-de-Software-UCC/backend/utils"

)

type usuarioServicio struct{
	usuarioCliente client.UsuarioClienteInterface
}

type usuarioServiceInterface interface {
	Login(email string, password string) (string , error)
	GetUsuariobyEmail(email string) (dominio.UsuarioData, e.ApiError)
	CrearUsuario(newusuario dominio.UsuarioData) (dominio.UsuarioData,e.ApiError)
}


var (
	UsuarioServicio usuarioServiceInterface
)

func initUsuarioService(usuarioCliente client.UsuarioClienteInterface) usuarioServiceInterface {
	service := new(usuarioServicio)
	service.usuarioCliente = usuarioCliente
	return service
}

func init(){
	UsuarioServicio = initUsuarioService(client.UsuarioCliente)
}

func generateHash(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func(s *usuarioServicio) Login(email string, password string) (string,error){

	if strings.TrimSpace(email) == ""{
		return "", errors.New("Debe ingresar un email.")
	}

	if strings.TrimSpace(password) == ""{
		return "", errors.New("Debe ingresar una contraseña.")
	}

	hash :=  generateHash(password)
	log.Println("Hash de la contraseña ingresada:", hash) 

	usuarioo, err :=  s.usuarioCliente.GetUsuariobyEmail(email)

	if err != nil{
		return "", fmt.Errorf("Hubo un error al buscar el usuario en la Base de Datos.")
	
	}

	if hash != usuarioo.Passwordhash {
		return "", fmt.Errorf("Contraseña incorrecta.")
	}

	token := hash;
	return token, nil
}

func (s *usuarioServicio) GetUsuariobyEmail(email string) (dominio.UsuarioData, e.ApiError) {

	var usuarioo dao.Usuario
	usuarioo, err := s.usuarioCliente.GetUsuariobyEmail(email)
	var us dominio.UsuarioData

	if err != nil {
		return us, e.NewBadRequestApiError("Usuario no encontrado")
	}
	
	us.Nombre = usuarioo.Nombre
	us.Apellido = usuarioo.Apellido
	us.Tipo = usuarioo.Tipo
	us.Email = usuarioo.Email
	us.Passwordhash = usuarioo.Passwordhash
	us.UsuarioID = usuarioo.UsuarioID
	

	return us, nil
}

func (s *usuarioServicio) CrearUsuario (newusuario dominio.UsuarioData) (dominio.UsuarioData,e.ApiError) { 

	var user dao.Usuario

	user.Nombre = newusuario.Nombre
	user.Apellido = newusuario.Apellido
	user.Email = newusuario.Email

	hash := generateHash(newusuario.Passwordhash)
	log.Println("Hash de la contraseña ingresada:", hash) 


	user.Passwordhash = hash
	newusuario.Passwordhash = user.Passwordhash
	user.Tipo = newusuario.Tipo

	user = s.usuarioCliente.CrearUsuario(user)

	newusuario.UsuarioID = user.UsuarioID

	return newusuario, nil

}

