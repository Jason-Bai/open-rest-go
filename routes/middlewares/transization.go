package middlewares

import (
	structs "github.com/Jason-Bai/open-rest-go/structs/middlewares"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Transization a database transization middleware
func Transization(db *gorm.DB, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := db.Begin()

		defer func() {
			if err := recover(); err != nil {
				entry := err.(*logrus.Entry)
				log.WithFields(logrus.Fields{
					"omg":         true,
					"err_animal":  entry.Data["animal"],
					"err_size":    entry.Data["size"],
					"err_level":   entry.Level,
					"err_message": entry.Message,
					"number":      100,
				}).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
				tx.Rollback()
			}
		}()

		rc, _ := structs.GetRequestContext(c)

		rc.SetTx(tx)

		c.Next()

		tx.Commit()
	}
}
