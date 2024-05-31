package usuario

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
