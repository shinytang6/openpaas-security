package models

import (
	"fmt"
	db "github.com/shinytang6/openpaas-security/backend/database"
	"strconv"

	"github.com/shinytang6/openpaas-security/utils"
)

type Experiment struct {
	Id    		 int    `json:"id" form:"id"`
	ExperimentId int 	`json:"experimentId" form:"experimentId"`
	Name  		 string `json:"name" form:"name"`
	Date  		 string `json:"date" form:"date"`
	PersonId     int 	`json:"personId" form:"personId"`
	Address    string
}

var PublicIP string = "192.168.99.100"

func (e *Experiment) AddExperiment() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO Experiment(experimentId, name) VALUES (?, ?)", e.ExperimentId, e.Name)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (e *Experiment) GetExperiments() (experiments []Experiment, err error) {
	experiments = make([]Experiment, 0)
	rows, err := db.SqlDB.Query("SELECT id, experimentId, name, date, personId FROM Experiment")
	defer rows.Close()

	if err != nil {
		return
	}
	fmt.Println(rows.Columns())
	for rows.Next() {
		var experiment Experiment
		rows.Scan(&experiment.Id, &experiment.ExperimentId, &experiment.Name, &experiment.Date, &experiment.PersonId)
		nodePort := utils.GetOneExperiment(experiment.Name).Spec.Ports[0].NodePort
		port := fmt.Sprint(nodePort)
		experiment.Address = PublicIP + ":" + port
		experiments = append(experiments, experiment)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}


func (e *Experiment) GetExperiment(id int, experimentId int, name string) (experiment Experiment, err error) {
	rows, err := db.SqlDB.Query("SELECT * FROM Experiment where id=? and experimentId=?", id, experimentId)
	defer rows.Close()

	if err != nil {
		return
	}

	rows.Next()
	rows.Scan(&experiment.Id, &experiment.ExperimentId, &experiment.Name)

	if err = rows.Err(); err != nil {
		return
	}
	return
}


func CreateExperiment(name string, config string, people int, date string) (err error) {
	// wtf??? QueryRow
	row := db.SqlDB.QueryRow("SELECT max(experimentId) FROM Experiment")
	var max int
	row.Scan(&max)
	for i:=0; i<people; i++ {
		utils.StartOneExperiment("", name+"-"+strconv.Itoa(i))
		_, err := db.SqlDB.Query("INSERT INTO Experiment(experimentId, config, personId, date, name) VALUES(?, ?, ?, ?, ?)", max+1, config, i, date, name+"-"+strconv.Itoa(i))
		if err != nil {
			return err
		}

	}
	return err
}

func DeleteExperiment(name string) (err error) {
	rows, err := db.SqlDB.Query("DELETE FROM Experiment where name=?", name)
	defer rows.Close()
	if err != nil {
		return
	}
	fmt.Println("tl")
	fmt.Println(name+"!!!!!!!!!!!!!!!!!!!!!")
	utils.DeleteOneExperiment(name)
	return err
}

func RestartExperiment(name string) (err error) {
	utils.RestartOneExperiment(name)
	return err
}