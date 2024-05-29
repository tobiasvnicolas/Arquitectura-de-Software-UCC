package controlador

import(
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"Arquitectura-de-Software-UCC/backend/dominio/cursos"
	"Arquitectura-de-Software-UCC/backend/servicios/cursos"
	"net/http"

)

func CrearCurso(c *gin.Context){
	var newcurso dominio.CursoData

	err := c.BindJSON(&newcurso)
	log.Debug(newcurso)

	if err != nil{
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newcurso, er := servicios.CursoServicio.CrearCurso(newcurso)
	if er != nil{
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, newcurso)

}