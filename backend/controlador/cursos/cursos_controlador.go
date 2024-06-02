package cursosControlador

import(
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"Arquitectura-de-Software-UCC/backend/dominio"
	"Arquitectura-de-Software-UCC/backend/servicios"
	"net/http"	
	"strings"
	"fmt"
	"strconv"

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

func GetCursoById (c *gin.Context){
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



func SearchCursos (c *gin.Context){
	palabra := strings.TrimSpace(c.Param("palabra"))

	results, err := servicios.CursoServicio.SearchCursos(palabra)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dominio.Resultado{
			Mensaje: fmt.Sprintf("Error al buscar: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, dominio.CursoBuscado{
		Results: results,
	})


}