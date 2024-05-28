package servicios

import(
	"Arquitectura-de-Software-UCC/backend/utils/errors"
	"Arquitectura-de-Software-UCC/backend/dominio/usuario"
	"Arquitectura-de-Software-UCC/backend/dao"
)

func Login(email string, password string) (string,error){

	if strings.TrimSpace(email) == ""{
		return "", errors.New("Debe ingresar un email.")
	}

	if strings.TrimSpace(password) == ""{
		return "", errors.New("Debe ingresar una contraseña.")
	}

	hash := fmt.Sprintf("%x", mdS.Sum([]byte(password)))

	Usuario, err := clients.GetUsuariobyEmail(email)
	if err != nil{
		return "", fmt.Errorf("Hubo un error al buscar el usuario en la Base de Datos.")
	
	}

	if hash != Usuario.Passwordhash {
		return "", fmt.Errorf("Contraseña incorrecta.")
	}

	token := hash;
	return token, nil
}

type usuarioService struct{}

type usuarioServiceInterface interface {
	GetUsuariobyEmail(email string) (dominio.UsuarioData, e.ApiError)
}

var (
	UsuarioService usuarioServiceInterface
)

func init() {
	UsuarioService = &usuarioService{}
}

func(s *usuarioService) GetUsuariobyEmail(email string) (dominio.UsuarioData, e.ApiError){
	var usuario dao.Usuario = clientes.GetUsuariobyEmail(email)
	var usuarioData dominio.UsuarioData

	usuarioData.Email = usuario.Email
	usuarioData.Nombre = usuarioData.Nombre
	
	return usuarioData, nil

}