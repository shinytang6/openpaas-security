package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shinytang6/openpaas-security/backend/models"
	"github.com/shinytang6/openpaas-security/backend/util"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddExperimentApi(c *gin.Context) {
	//c.Request.ParseForm()
	//for k, v := range c.Request.PostForm {
	//	fmt.Printf("k:%v\n", k)
	//	fmt.Printf("v:%v\n", v)
	//}
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	config := c.PostForm("config")
	people, _ := strconv.Atoi(c.PostForm("people"))
	date := c.PostForm("date")

	guide, err := c.FormFile("guide_path")
	if err != nil {
		c.String(http.StatusBadRequest, "a bad request")
	}
	filename := guide.Filename

	fileMeta := models.FileMeta{
		FileName: filename,
		Location: "./tempFiles/" + guide.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		FileSize: guide.Size,
	}

	//if err := c.SaveUploadedFile(guide, "./tempFiles/" + guide.Filename); err != nil {
	//	c.String(http.StatusBadRequest, fmt.Sprintf("upload file err : %s", err.Error()))
	//	return
	//}

	// 原生创建法
	newFile, err := os.Create("./tempFiles/"+filename)
	if err != nil {
		fmt.Println("Failed to create file, err:%s\n", err.Error())
		return
	}
	defer newFile.Close()

	file, _ := guide.Open()
	_, error := io.Copy(newFile, file)

	if error != nil {
		fmt.Println("Failed to save data into file, err: %s\n", err.Error())
		return
	}

	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)

	models.UpdateFileMeta(fileMeta)

	err = models.CreateExperiment(name, config, people, date, fileMeta.FileSha1, filename, desc)
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successfully")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})

	//experimentId := c.Request.FormValue("experimentId")
	//name := c.Request.FormValue("name")
	//experimentId, _ := strconv.Atoi(c.PostForm("experimentId"))
	//name := c.PostForm("name")
	//
	//e := models.Experiment{ExperimentId: experimentId, Name: name}
	//
	//ra, err := e.AddExperiment()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//msg := fmt.Sprintf("insert successful %d", ra)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": msg,
	//})
}


func GetAllExperimentsApi(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	e := models.Experiment{}

	experiments, err := e.GetExperiments()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful result: %+v", experiments)
	c.JSON(http.StatusOK, gin.H{
		"data": experiments,
		"msg": msg,
	})
}

func GetExperimentsByPersonId(c *gin.Context) {
	personId, _ := strconv.Atoi(c.Query("id"))
	//experimentId, _ := strconv.Atoi(c.Query("experimentId"))
	//name := c.Query("name")

	e := models.Experiment{}

	experiments, err := e.GetExperimentsByPersonId(personId)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("get successful result: %+v", experiments)
	c.JSON(http.StatusOK, gin.H{
		"data": experiments,
		"msg": msg,
	})
}

func CreateExperiment(c *gin.Context) {
	name := c.Query("name")
	config := c.Query("config")
	people, _ := strconv.Atoi(c.Query("people"))
	date := c.Query("date")

	err := models.CreateExperiment(name, config, people, date, "1", "1", "1")
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("create successful")
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": msg,
	})
}

func DeleteExperiment(c *gin.Context) {
	name := c.Query("name")
	err := models.DeleteExperiment(name)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successfully")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func RestartExperiment(c *gin.Context) {
	name := c.Query("name")
	err := models.RestartExperiment(name)
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("restart successfully")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func TestExperiment(c *gin.Context) {
	//c.Header("content-disposition", `attachment; filename=` + "chat.py")
	//c.Writer.Header().Add("Content-Type", "application/octet-stream")
	filehash := c.Query("filehash")
	fMeta := models.GetFileMeta(filehash)
	c.JSON(http.StatusOK, gin.H{
		//"data": experiment,
		"msg": fMeta,
	})
}

func DownloadFile(c *gin.Context) {
	filehash := c.Query("filehash")
	fMeta := models.GetFileMeta(filehash)

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fMeta.FileName))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")

	c.File(fMeta.Location)
}