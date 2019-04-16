package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/shinytang6/openpaas-security/backend/models"
	"strconv"
)

func GetStudentApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	experimentId, _ := strconv.Atoi(c.Query("experimentId"))
	name := c.Query("name")

	e := models.Experiment{Id: id, ExperimentId: experimentId, Name: name}

	experiment, err := e.GetExperiment(id, experimentId, name)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", experiment)
	c.JSON(http.StatusOK, gin.H{
		"data": experiment,
		"msg": msg,
	})
}

func LoginStudentApi(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	s := models.Student{}

	student, err := s.GetStudent(name, password)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", student)
	if student.Name == "" {
		msg = fmt.Sprintf("empty result")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": msg,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
		"msg": msg,
	})
}