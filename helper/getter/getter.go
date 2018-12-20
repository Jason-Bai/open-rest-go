package getter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Getter get a model instance with special gorm model and key
func Getter(model *gorm.Model, keyPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("getter before")
		c.Next()
	}
}
