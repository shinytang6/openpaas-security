package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shinytang6/openpaas-security/backend/models"
	"log"
	"net/http"
)

func GetAllSysAdminsApi(c *gin.Context) {
	s := models.SysAdmin{}

	sysAdmins, err := s.GetSysAdmins()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", sysAdmins)
	c.JSON(http.StatusOK, gin.H{
		"data": sysAdmins,
		"msg": msg,
	})
}

func LoginSysAdminApi(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	s := models.SysAdmin{}

	sysAdmin, err := s.GetSysAdmin(name, password)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", sysAdmin)
	if sysAdmin.Name == "" {
		msg = fmt.Sprintf("empty result")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": msg,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": sysAdmin,
		"msg": msg,
	})
}

func DeleteSysAdmin(c *gin.Context) {
	name := c.Query("name")
	err := models.DeleteSysAdmin(name)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}

func CreateSysAdmin(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	sysAdminId := c.Query("sysAdminId")
	email := c.Query("email")
	phone := c.Query("phone")

	err := models.CreateSysAdmin(name, password, sysAdminId, email, phone)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("create successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}

func UpdateSysAdminApi(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	sysAdminId := c.Query("sysAdminId")
	email := c.Query("email")
	phone := c.Query("phone")

	s := models.SysAdmin{}

	err := s.UpdateSysAdmin(name, password, sysAdminId, email, phone)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("update successful")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}