package inscripcion

import (
	"Arquitectura-de-Software-UCC/backend/dominio"
	"Arquitectura-de-Software-UCC/backend/servicios"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CrearInscripcion(c *gin.Context) {
	var newinscripcion dominio.InscripcionData

	err := c.BindJSON(&newinscripcion)

	log.Debug(newinscripcion)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newinscripcion, er := servicios.InscripcionServicio.CrearInscripcion(newinscripcion)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, newinscripcion)
}

func GetInscripcionByUserId(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cursoData, er := servicios.InscripcionServicio.GetInscripcionByUserId(id)

	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cursos": cursoData})

}

/*
func GetCursosByUsuarioid(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//Hasta aca solo verifico que no haya errores con el id del usuario

	cursoData, er := servicios.InscripcionServicio.GetInscripcionByUserId(id)


	c.JSON(http.StatusOK, cursoData)
}
*/
