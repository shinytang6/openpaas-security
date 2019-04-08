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

	return router
}