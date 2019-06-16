package models

import (
	db "github.com/shinytang6/openpaas-security/backend/database"
	"github.com/shinytang6/openpaas-security/backend/util"
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
	rows, err := db.SqlDB.Query("SELECT * FROM UserAdmin where name=? and password=?", name, util.MD5(password))
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

func (u *UserAdmin) UpdateUserAdmin(name string, password string, email string, phone string) error {
	if name == "" {
		name = u.Name
	}
	if password == "" {
		password = u.Password
	}

	if email == "" {
		email = u.Email
	}
	if phone == "" {
		phone = u.Phone
	}
	_, err := db.SqlDB.Exec("UPDATE UserAdmin SET name=?, password=?, email=?, phone=? where userId=?", name, util.MD5(password), email, phone, u.UserId)

	if err != nil {
		return err
	}

	return nil
}

func GetUserAdminById(id int) (userAdmin UserAdmin, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM UserAdmin where userId=?", id)
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