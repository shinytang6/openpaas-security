package models

import (
	db "github.com/shinytang6/openpaas-security/backend/database"
)

type Teacher struct {
	UserId       int        `json:"userId" form:"userId"`
	Password     string     `json:"password" form:"password"`
	Name  		 string     `json:"name" form:"name"`
	TeacherId    string     `json:"teacherId" form:"teacherId"`
	Email        string 	`json:"email" form:"email"`
	Phone        string 	`json:"phone" form:"phone"`
}


func (t *Teacher) GetTeacher(name string, password string) (teacher Teacher, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM Teacher where name=? and password=?", name, password)
	defer rows.Close()

	if err != nil {
		return
	}

	rows.Next()
	rows.Scan(&teacher.UserId, &teacher.Password, &teacher.Name, &teacher.TeacherId, &teacher.Email, &teacher.Phone)

	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (t *Teacher) GetTeachers() (teachers []Teacher, err error) {
	teachers = make([]Teacher, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM Teacher")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var teacher Teacher
		rows.Scan(&teacher.UserId, &teacher.Password, &teacher.Name, &teacher.TeacherId, &teacher.Email, &teacher.Phone)
		teachers = append(teachers, teacher)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (t *Teacher) UpdateTeacher(name string, password string, teacherId string, email string, phone string) error {
	if name == "" {
		name = t.Name
	}
	if password == "" {
		password = t.Password
	}
	if teacherId == "" {
		teacherId = t.TeacherId
	}
	if email == "" {
		email = t.Email
	}
	if phone == "" {
		phone = t.Phone
	}

	_, err := db.SqlDB.Exec("UPDATE Teacher SET name=?, password=?, TeacherId=?, email=?, phone=?", name, password, teacherId, email, phone)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTeacher(name string) (err error) {
	rows, err := db.SqlDB.Query("DELETE FROM Teacher where name=?", name)
	defer rows.Close()
	if err != nil {
		return
	}
	return err
}

func CreateTeacher(name string, password string, teacherId string, email string, phone string) error {
	_, err := db.SqlDB.Exec("INSERT INTO Teacher (name, password, teacherId, email, phone) VALUES(?, ?, ?, ?, ?)", name, password, teacherId, email, phone)

	if err != nil {
		return err
	}

	return nil
}


func GetTeacherByName(name string) (teacher Teacher, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM Teacher where name=?", name)
	defer rows.Close()

	if err != nil {
		return
	}

	rows.Next()
	rows.Scan(&teacher.UserId, &teacher.Password, &teacher.Name, &teacher.TeacherId, &teacher.Email, &teacher.Phone)

	if err = rows.Err(); err != nil {
		return
	}
	return
}

func GetTeacherById(id int) (teacher Teacher, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM Teacher where userId=?", id)
	defer rows.Close()

	if err != nil {
		return
	}

	rows.Next()
	rows.Scan(&teacher.UserId, &teacher.Password, &teacher.Name, &teacher.TeacherId, &teacher.Email, &teacher.Phone)

	if err = rows.Err(); err != nil {
		return
	}
	return
}