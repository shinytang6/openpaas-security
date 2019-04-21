package models

import (
	db "github.com/shinytang6/openpaas-security/backend/database"
)

type SysAdmin struct {
	UserId         int        `json:"userId" form:"userId"`
	Password       string     `json:"password" form:"password"`
	Name  		   string     `json:"name" form:"name"`
	SystemAdminId  string     `json:"systemAdminId" form:"systemAdminId"`
	Email          string 	  `json:"email" form:"email"`
	Phone          string 	  `json:"phone" form:"phone"`
}


func (s *SysAdmin) GetSysAdmin(name string, password string) (sysAdmin SysAdmin, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM SystemAdmin where name=? and password=?", name, password)
	defer rows.Close()

	if err != nil {
		return
	}

	rows.Next()
	rows.Scan(&sysAdmin.UserId, &sysAdmin.Password, &sysAdmin.Name, &sysAdmin.SystemAdminId, &sysAdmin.Email, &sysAdmin.Phone)

	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (s *SysAdmin) GetSysAdmins() (sysAdmins []SysAdmin, err error) {
	sysAdmins = make([]SysAdmin, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM SystemAdmin")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var sysAdmin SysAdmin
		rows.Scan(&sysAdmin.UserId, &sysAdmin.Password, &sysAdmin.Name, &sysAdmin.SystemAdminId, &sysAdmin.Email, &sysAdmin.Phone)
		sysAdmins = append(sysAdmins, sysAdmin)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func DeleteSysAdmin(name string) (err error) {
	rows, err := db.SqlDB.Query("DELETE FROM SystemAdmin where name=?", name)
	defer rows.Close()
	if err != nil {
		return
	}
	return err
}

func CreateSysAdmin(name string, password string, sysAdminId string, email string, phone string) error {
	_, err := db.SqlDB.Exec("INSERT INTO SystemAdmin (name, password, systemAdminId, email, phone) VALUES(?, ?, ?, ?, ?)", name, password, sysAdminId, email, phone)

	if err != nil {
		return err
	}

	return nil
}

func (s *SysAdmin) UpdateSysAdmin(name string, password string, sysAdminId string, email string, phone string) error {
	_, err := db.SqlDB.Exec("UPDATE SystemAdmin SET name=?, password=?, systemAdminId=?, email=?, phone=? where name=?", name, password, sysAdminId, email, phone, name)

	if err != nil {
		return err
	}

	return nil
}