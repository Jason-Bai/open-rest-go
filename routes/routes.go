package routes

import (
	"fmt"

	"github.com/Jason-Bai/open-rest-go/apis/user/delivery/http"
	"github.com/Jason-Bai/open-rest-go/routes/middlewares"
	structs "github.com/Jason-Bai/open-rest-go/structs/middlewares"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
}

func setMiddlewares(db *gorm.DB, log *logrus.Logger) {
	// global middlewares
	globalMiddlewares := []gin.HandlerFunc{
		middlewares.Logger(log),
		middlewares.Context(),
		middlewares.Transization(db, log),
	}

	engine.Use(globalMiddlewares...)
}

// Routes add
func Routes(db *gorm.DB, log *logrus.Logger) *gin.Engine {
	// 1. set global middlewares
	setMiddlewares(db, log)

	// 2. set public routes
	engine.GET("/ping", func(c *gin.Context) {
		rc, _ := structs.GetRequestContext(c)

		fmt.Println("x-request-id:", rc.RequestID())

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 3. set api routes
	http.NewUserHttpRoutes(engine)

	return engine
}
