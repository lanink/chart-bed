package routers

import (
	"chart-bed/src/controllers"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	fileCtlRegister(engine)
}

func fileCtlRegister(engine *gin.Engine) {
	ftlCtl := controllers.FileController{}
	engine.POST("/upload", ftlCtl.Upload)
	engine.GET("/images/:image", ftlCtl.Image)
}
