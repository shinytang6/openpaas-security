package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shinytang6/openpaas-security/backend/models"
	"log"
	"net/http"
)


func GetAllTeachersApi(c *gin.Context) {
	t := models.Teacher{}

	teachers, err := t.GetTeachers()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", teachers)
	c.JSON(http.StatusOK, gin.H{
		"data": teachers,
		"msg": msg,
	})
}

func UpdateTeacherApi(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	teacherId := c.Query("teacherId")
	email := c.Query("email")
	phone := c.Query("phone")

	t := models.Teacher{}
	fmt.Println("teacher test")
	fmt.Println(name, password, teacherId, email, phone)
	err := t.UpdateTeacher(name, password, teacherId, email, phone)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("update successful")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeletTeacher(c *gin.Context) {
	name := c.Query("name")
	err := models.DeleteTeacher(name)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}

func CreateTeacher(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	teacherId := c.Query("teacherId")
	email := c.Query("email")
	phone := c.Query("phone")

	err := models.CreateTeacher(name, password, teacherId, email, phone)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("create successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}

func LoginTeacherApi(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	t := models.Teacher{}

	teacher, err := t.GetTeacher(name, password)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", teacher)
	if teacher.Name == "" {
		msg = fmt.Sprintf("empty result")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": msg,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": teacher,
		"msg": msg,
	})
}