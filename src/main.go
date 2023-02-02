package main

import "C"
import (
	"chart-bed/src/configs"
	"chart-bed/src/routers"
	"chart-bed/src/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	config := configs.ReadConfig()

	_, exist := os.LookupEnv("APP_CHAR_BED_DEVELOP")

	if exist {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	routers.Register(engine)

	_, err := utils.CheckAndCreateDir(config.Dir())

	if err != nil {
		log.Fatalln(err)
	}

	_ = engine.Run(config.Server())
}
