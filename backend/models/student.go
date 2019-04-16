package models

import (
	db "github.com/shinytang6/openpaas-security/backend/database"
)

type Student struct {
	UserId       int        `json:"userId" form:"userId"`
	Password     string     `json:"password" form:"password"`
	Name  		 string     `json:"name" form:"name"`
	StudentId    string     `json:"studentId" form:"studentId"`
	Class        string 	`json:"class" form:"class"`
	Email        string 	`json:"email" form:"email"`
	Phone        string 	`json:"phone" form:"phone"`
}


func (s *Student) GetStudent(name string, password string) (student Student, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM Student where name=? and password=?", name, password)
	defer rows.Close()

	if err != nil {
		return
	}

	rows.Next()
	rows.Scan(&student.UserId, &student.Password, &student.Name, &student.StudentId, &student.Class, &student.Email, &student.Phone)

	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (s *Student) GetStudents() (students []Student, err error) {
	students = make([]Student, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM Student")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var student Student
		rows.Scan(&student.UserId, &student.Password, &student.Name, &student.StudentId, &student.Class, &student.Email, &student.Phone)
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
