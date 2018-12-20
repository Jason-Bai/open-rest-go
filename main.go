package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Jason-Bai/open-rest-go/config"
	"github.com/Jason-Bai/open-rest-go/logger"
	"github.com/Jason-Bai/open-rest-go/models"
	"github.com/Jason-Bai/open-rest-go/routes"
)

func main() {
	// 1. load configuration files
	if err := config.LoadConfig("./config/files"); err != nil {
		panic(fmt.Errorf("load configuration failed: %s", err))
	}

	// 2. create a logger
	log := logger.NewLogger()

	// 2. connect to mysql database
	db := models.OpenDB(log)
	defer db.Close()

	// 3. wire up api routings
	r := routes.Routes(db, log)

	addr := fmt.Sprintf(":%v", config.Config.SERVICE.PORT)

	if "release" == gin.Mode() {
		log.Infof("server is started at %v port", config.Config.SERVICE.PORT)
	}

	// 4. listen to 8080 port
	panic(r.Run(addr)) // listen and serve on 0.0.0.0:8080
}
