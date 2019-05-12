package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/shinytang6/openpaas-security/backend/models"
	"strconv"
)


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

func GetAllStudentsApi(c *gin.Context) {
	s := models.Student{}

	students, err := s.GetStudents()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", students)
	c.JSON(http.StatusOK, gin.H{
		"data": students,
		"msg": msg,
	})
}

func UpdateStudentApi(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	name := c.Query("name")
	password := c.Query("password")
	studentId := c.Query("studentId")
	class := c.Query("class")
	email := c.Query("email")
	phone := c.Query("phone")

	s, _ := models.GetStudentById(id)
	err := s.UpdateStudent(name, password, studentId, class, email, phone)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("update successful")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeleteStudent(c *gin.Context) {
	name := c.Query("name")
	err := models.DeleteStudent(name)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}

func CreateStudent(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	studentId := c.Query("studentId")
	class := c.Query("class")
	email := c.Query("email")
	phone := c.Query("phone")

	err := models.CreateStudent(name, password, studentId, class, email, phone)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("create successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}