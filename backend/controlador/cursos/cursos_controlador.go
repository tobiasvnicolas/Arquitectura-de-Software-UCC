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
	id := c.Param("id")
	cursoID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}


	cursoData, er := servicios.CursoServicio.GetCursoById(cursoID)

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

func GetCursosByIds (c *gin.Context){

	var cursoids []int64

	for _, idStr := range c.QueryArray("ids"){
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso no válido"})
			return
		}
		cursoids = append(cursoids, id)
	}



	cursoData, er := servicios.CursoServicio.GetCursosByIds(cursoids)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, cursoData)

}