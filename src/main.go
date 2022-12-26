package main

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	imageflow "github.com/imazen/imageflow-go"
	"log"
	"net/http"
	"strconv"
)

func main() {

	// w=300&h=300     =>  is ok
	// dpr mode, scale, anchor, sflip
	// more params  https://docs.imageflow.io/introduction.html

	step := imageflow.NewStep()

	// for release
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		dist := "./images/" + file.Filename
		// save file to dist
		err := c.SaveUploadedFile(file, dist)

		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{"url": fmt.Sprintf("http://%s/images/%s",
			c.Request.Host, file.Filename),
			"code": 0})
	})

	r.GET("/images/:image", func(c *gin.Context) {
		image := c.Param("image")
		//width=100&height=100&mode=max&scale=down
		width, _ := strconv.ParseFloat(c.Query("w"), 64)
		height, _ := strconv.ParseFloat(c.Query("h"), 64)

		img, _ := step.Decode(imageflow.NewFile("./images/"+image)).
			ConstrainWithin(width, height).
			Encode(imageflow.GetBuffer("buf"), imageflow.MozJPEG{}).
			Execute()

		_, _ = c.Writer.Write(img["buf"])
	})

	_ = r.Run() // listen and serve on 0.0.0.0:8080

}
