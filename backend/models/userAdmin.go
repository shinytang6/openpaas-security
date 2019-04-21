package models

import (
	db "github.com/shinytang6/openpaas-security/backend/database"
)

type UserAdmin struct {
	UserId         int        `json:"userId" form:"userId"`
	Password       string     `json:"password" form:"password"`
	Name  		   string     `json:"name" form:"name"`
	UserAdminId    string     `json:"userAdminId" form:"userAdminId"`
	Email          string 	  `json:"email" form:"email"`
	Phone          string 	  `json:"phone" form:"phone"`
}


func (u *UserAdmin) GetUserAdmin(name string, password string) (userAdmin UserAdmin, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM UserAdmin where name=? and password=?", name, password)
	defer rows.Close()

	if err != nil {
		return
	}

	rows.Next()
	rows.Scan(&userAdmin.UserId, &userAdmin.Password, &userAdmin.Name, &userAdmin.UserAdminId, &userAdmin.Email, &userAdmin.Phone)

	if err = rows.Err(); err != nil {
		return
	}
	return
}