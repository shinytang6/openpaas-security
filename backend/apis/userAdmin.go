package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shinytang6/openpaas-security/backend/models"
	"log"
	"net/http"
)


func LoginUserAdminApi(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	u := models.UserAdmin{}

	userAdmin, err := u.GetUserAdmin(name, password)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful %d", userAdmin)
	if userAdmin.Name == "" {
		msg = fmt.Sprintf("empty result")
		c.JSON(http.StatusNotFound, gin.H{
			"msg": msg,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userAdmin,
		"msg": msg,
	})
}