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
	FileHash     string `json:"fileHash" form:"fileHash"`
	FileName     string `json:"fileName" form:"fileName"`
	Description  string `json:"description" form:"description"`
	Address      string
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
	rows, err := db.SqlDB.Query("SELECT id, experimentId, name, date, personId, fileHash, fileName, description FROM Experiment")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var experiment Experiment
		rows.Scan(&experiment.Id, &experiment.ExperimentId, &experiment.Name, &experiment.Date, &experiment.PersonId, &experiment.FileHash, &experiment.FileName, &experiment.Description)
		if len(utils.GetOneExperiment(experiment.Name).Spec.Ports) > 0 {
			nodePort := utils.GetOneExperiment(experiment.Name).Spec.Ports[0].NodePort
			port := fmt.Sprint(nodePort)
			experiment.Address = PublicIP + ":" + port
		}
		experiments = append(experiments, experiment)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}


func (e *Experiment) GetExperimentsByPersonId(personId int) (experiments []Experiment, err error) {
	experiments = make([]Experiment, 0)
	rows, err := db.SqlDB.Query("SELECT id, experimentId, name, date, personId, fileHash, fileName, description FROM Experiment where personId=?", personId)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var experiment Experiment
		rows.Scan(&experiment.Id, &experiment.ExperimentId, &experiment.Name, &experiment.Date, &experiment.PersonId, &experiment.FileHash, &experiment.FileName, &experiment.Description)
		if len(utils.GetOneExperiment(experiment.Name).Spec.Ports) > 0 {
			nodePort := utils.GetOneExperiment(experiment.Name).Spec.Ports[0].NodePort
			port := fmt.Sprint(nodePort)
			experiment.Address = PublicIP + ":" + port
		}
		experiments = append(experiments, experiment)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}


func CreateExperiment(name string, config string, people int, date string, fileHash string, fileName string, desc string) (err error) {
	// wtf??? QueryRow
	row := db.SqlDB.QueryRow("SELECT max(experimentId) FROM Experiment")
	var max int
	row.Scan(&max)
	for i:=0; i<people; i++ {
		utils.StartOneExperiment("", name+"-"+strconv.Itoa(i+1))
		_, err := db.SqlDB.Query("INSERT INTO Experiment(experimentId, config, personId, date, name, fileHash, fileName, description) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", max+1, config, i+1, date, name+"-"+strconv.Itoa(i+1), fileHash, fileName, desc)
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

	utils.DeleteOneExperiment(name)
	return err
}

func RestartExperiment(name string) (err error) {
	utils.RestartOneExperiment(name)
	return err
}