package controllers

import (
	"chart-bed/src/configs"
	"chart-bed/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imazen/imageflow-go"
	"log"
	"net/http"
	"strconv"
)

type FileController struct {
}

func (fileCtl *FileController) Upload(cxt *gin.Context) {
	file, _ := cxt.FormFile("file")

	dir := configs.ReadConfig().Dir()

	exist, err := utils.CheckAndCreateDir(dir)

	if exist {
		dist := dir + file.Filename
		err := cxt.SaveUploadedFile(file, dist)
		if err != nil {
			log.Println(err)
		}
		cxt.JSON(http.StatusOK, gin.H{"url": fmt.Sprintf("http://%s/images/%s",
			cxt.Request.Host, file.Filename),
			"code": 0})
	} else {
		cxt.JSON(http.StatusServiceUnavailable, gin.H{"err": err})
	}
}

func (fileCtl *FileController) Image(cxt *gin.Context) {
	image := cxt.Param("image")
	width, _ := strconv.ParseFloat(cxt.Query("w"), 64)
	height, _ := strconv.ParseFloat(cxt.Query("h"), 64)
	step := imageflow.NewStep()
	img, _ := step.Decode(imageflow.NewFile("./images/"+image)).
		ConstrainWithin(width, height).
		Encode(imageflow.GetBuffer("buf"), imageflow.MozJPEG{}).
		Execute()
	_, _ = cxt.Writer.Write(img["buf"])
}
