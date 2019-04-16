package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shinytang6/openpaas-security/backend/apis"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.IndexApi)

	router.POST("/api/experiment/add", apis.AddExperimentApi)

	router.GET("/api/experiment/getall", apis.GetAllExperimentsApi)

	router.GET("/api/experiment/get", apis.GetExperimentApi)

	router.GET("/api/experiment/add", apis.CreateExperiment)

	router.GET("/api/experiment/delete", apis.DeleteExperiment)


	router.GET("/api/student/get", apis.GetStudentApi)

	router.GET("/api/student/login", apis.LoginStudentApi)

	router.GET("/api/student/getall", apis.GetAllStudentsApi)
	return router
}