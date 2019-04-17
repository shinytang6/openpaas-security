package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/shinytang6/openpaas-security/backend/models"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddExperimentApi(c *gin.Context) {
	//experimentId := c.Request.FormValue("experimentId")
	//name := c.Request.FormValue("name")
	experimentId, _ := strconv.Atoi(c.PostForm("experimentId"))
	name := c.PostForm("name")

	e := models.Experiment{ExperimentId: experimentId, Name: name}

	ra, err := e.AddExperiment()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}


func GetAllExperimentsApi(c *gin.Context) {
	e := models.Experiment{}

	experiments, err := e.GetExperiments()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", experiments)
	c.JSON(http.StatusOK, gin.H{
		"data": experiments,
		"msg": msg,
	})
}

func GetExperimentApi(c *gin.Context) {
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

func CreateExperiment(c *gin.Context) {
	name := c.Query("name")
	people, _ := strconv.Atoi(c.Query("people"))
	date := c.Query("date")

	err := models.CreateExperiment(name, people, date)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("create successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}

func DeleteExperiment(c *gin.Context) {
	name := c.Query("name")
	err := models.DeleteExperiment(name)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}