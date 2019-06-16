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

func (s *Student) UpdateStudent(name string, password string, studentId string, class string, email string, phone string) error {
	if name == "" {
		name = s.Name
	}
	if password == "" {
		password = s.Password
	}
	if studentId == "" {
		studentId = s.StudentId
	}
	if class == "" {
		class = s.Class
	}
	if email == "" {
		email = s.Email
	}
	if phone == "" {
		phone = s.Phone
	}

	_, err := db.SqlDB.Exec("UPDATE Student SET name=?, password=?, studentId=?, class=?, email=?, phone=? where userId=?", name, password, studentId, class, email, phone, s.UserId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteStudent(name string) (err error) {
	rows, err := db.SqlDB.Query("DELETE FROM Student where name=?", name)
	defer rows.Close()
	if err != nil {
		return
	}
	return err
}

func CreateStudent(name string, password string, studentId string, class string, email string, phone string) error {
	_, err := db.SqlDB.Exec("INSERT INTO Student (name, password, studentId, class, email, phone) VALUES(?, ?, ?, ?, ?, ?)", name, password, studentId, class, email, phone)

	if err != nil {
		return err
	}

	return nil
}

func GetStudentByName(name string) (student Student, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM Student where name=?", name)
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

func GetStudentById(id int) (student Student, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM Student where userId=?", id)
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