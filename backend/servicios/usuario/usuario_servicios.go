package servicios

import(
	client "Arquitectura-de-Software-UCC/backend/clientes/usuario"
	"Arquitectura-de-Software-UCC/backend/dao"
	"Arquitectura-de-Software-UCC/backend/dominio/usuario"
	"fmt"
	"errors"
	"strings"
	"crypto/md5"
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


func(s *usuarioServicio) Login(email string, password string) (string,error){

	var usuario dao.Usuario
	if strings.TrimSpace(email) == ""{
		return "", errors.New("Debe ingresar un email.")
	}

	if strings.TrimSpace(password) == ""{
		return "", errors.New("Debe ingresar una contraseña.")
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))

	usuario, err :=  s.usuarioCliente.GetUsuariobyEmail(email)

	if err != nil{
		return "", fmt.Errorf("Hubo un error al buscar el usuario en la Base de Datos.")
	
	}

	if hash != usuario.Passwordhash {
		return "", fmt.Errorf("Contraseña incorrecta.")
	}

	token := hash;
	return token, nil
}

func (s *usuarioServicio) GetUsuariobyEmail(email string) (dominio.UsuarioData, e.ApiError) {

	var usuario dao.Usuario
	usuario, err := s.usuarioCliente.GetUsuariobyEmail(email)
	var us dominio.UsuarioData

	if err != nil {
		return us, e.NewBadRequestApiError("usuario not found")
	}
	us.Nombre = usuario.Nombre
	us.Apellido = usuario.Apellido
	us.Tipo = usuario.Tipo
	us.Email = usuario.Email
	us.Passwordhash = usuario.Passwordhash
	us.UsuarioID = usuario.UsuarioID
	

	return us, nil
}

func (s *usuarioServicio) CrearUsuario (newusuario dominio.UsuarioData) (dominio.UsuarioData,e.ApiError) { 

	var user dao.Usuario

	user.Nombre = newusuario.Nombre
	user.Apellido = newusuario.Apellido
	user.Email = newusuario.Email

	hash := fmt.Sprintf("%x", md5.Sum([]byte(newusuario.Passwordhash)))

	user.Passwordhash = hash
	newusuario.Passwordhash = user.Passwordhash
	user.Tipo = newusuario.Tipo

	user = s.usuarioCliente.CrearUsuario(user)

	newusuario.UsuarioID = user.UsuarioID

	return newusuario, nil

}

