package controlador

import (
	"Arquitectura-de-Software-UCC/backend/dominio/cursos"
	"Arquitectura-de-Software-UCC/backend/servicios/cursos"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func CrearCurso(c *gin.Context) {
	var newcurso dominio.CursoData

	err := c.BindJSON(&newcurso)
	log.Debug(newcurso)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newcurso, er := servicios.CursoServicio.CrearCurso(newcurso)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, newcurso)

}

func GetCursoById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var cursoData dominio.CursoData

	cursoData, er := servicios.CursoServicio.GetCursoById(id)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, cursoData)

}
