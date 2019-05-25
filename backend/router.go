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

	router.GET("/api/experiment/get", apis.GetExperimentsByPersonId)

	router.GET("/api/experiment/add", apis.CreateExperiment)

	router.GET("/api/experiment/delete", apis.DeleteExperiment)

	router.GET("/api/experiment/restart", apis.RestartExperiment)


	//router.GET("/api/student/get", apis.GetStudentApi)

	router.GET("/api/student/login", apis.LoginStudentApi)

	router.GET("/api/student/getall", apis.GetAllStudentsApi)

	router.GET("/api/student/update", apis.UpdateStudentApi)

	router.GET("/api/student/delete", apis.DeleteStudent)

	router.GET("/api/student/add", apis.CreateStudent)


	router.GET("/api/teacher/getall", apis.GetAllTeachersApi)
	router.GET("/api/teacher/update", apis.UpdateTeacherApi)
	router.GET("/api/teacher/delete", apis.DeleteTeacher)
	router.GET("/api/teacher/add", apis.CreateTeacher)
	router.GET("/api/teacher/login", apis.LoginTeacherApi)

	router.GET("/api/sysAdmin/getall", apis.GetAllSysAdminsApi)
	router.GET("/api/sysAdmin/delete", apis.DeleteSysAdmin)
	router.GET("/api/sysAdmin/update", apis.UpdateSysAdminApi)
	router.GET("/api/sysAdmin/add", apis.CreateSysAdmin)
	router.GET("/api/sysAdmin/login", apis.LoginSysAdminApi)

	router.GET("/api/userAdmin/login", apis.LoginUserAdminApi)
	router.GET("/api/userAdmin/update", apis.UpdateUserAdminApi)
	return router
}