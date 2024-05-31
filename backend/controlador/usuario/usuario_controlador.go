package usuario

import (
	"Arquitectura-de-Software-UCC/backend/dominio/usuario"
	"Arquitectura-de-Software-UCC/backend/servicios/usuario"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Login(c *gin.Context) {
	var request dominio.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dominio.Resultado{
			Mensaje: fmt.Sprintf("Request invalido: %s", err.Error()),
		})
		return
	}

	token, err := servicios.UsuarioServicio.Login(request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dominio.Resultado{
			Mensaje: fmt.Sprintf("Login no autorizado: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dominio.LoginResponse{
		Token: token,
	})
}

func GetUsuariobyEmail(c *gin.Context) {

	email := c.Param("email")
	var usuarioData dominio.UsuarioData

	usuarioData, err := servicios.UsuarioServicio.GetUsuariobyEmail(email)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usuarioData)

}

func CrearUsuario(c *gin.Context) {
	var newusuario dominio.UsuarioData

	err := c.BindJSON(&newusuario)

	log.Debug(newusuario)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newusuario, er := servicios.UsuarioServicio.CrearUsuario(newusuario)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, newusuario)
}

/*
import (
	usersDomain "Arquitectura-de-Software-UCC/backend/dominio/usuario"
	usersService "Arquitectura-de-Software-UCC/backend/servicios/usuario"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var loginRequest usersDomain.LoginRequest
	context.BindJSON(&loginRequest)
	response := usersService.Login(loginRequest)
	context.JSON(http.StatusOK, response)
}
*/
