package controlador

import(
	"github.com/gin-gonic/gin"
	"Arquitectura-de-Software-UCC/backend/dominio/usuario"
	"net/http"
	servicio "Arquitectura-de-Software-UCC/backend/servicios/usuario"

)

func Login(c *gin.Context){
	var request dominio.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, dominio.Resultado{
			Mensaje: fmt,Sprintf("Request invalido: %s", err.Error()),
		})
		return
	}

	token, err := servicios.Login(request.Email, request.Password)
	if err != nil{
		c.JSON(http.StatusUnathoride, dominio.Resultado{
			Mensaje: fmt.Sprintf("Login no autorizado: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dominio.LoginResponse{
		Token: token,
	})
}

func GetUsuariobyEmail(c *gin.Context){

email, _ := strconv.Atoi(c.Param("email"))
var usuarioData dominio.UsuarioData

usuarioData, err := servicio.usuario.GetUsuariobyEmail(email)

if err != nil{
	c.JSON(err.Status(),err)
	return
}

c.JSON(http.StatusOK, usuarioData)

}