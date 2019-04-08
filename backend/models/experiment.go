package models

import (
	db "github.com/shinytang6/openpaas-security/backend/database"
)

type Experiment struct {
	Id    		 int    `json:"id" form:"id"`
	ExperimentId int 	`json:"experimentId" form:"id"`
	Name  		 string `json:"name" form:"name"`
}

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
	rows, err := db.SqlDB.Query("SELECT id, experimentId, name FROM Experiment")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var experiment Experiment
		rows.Scan(&experiment.Id, &experiment.ExperimentId, &experiment.Name)
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