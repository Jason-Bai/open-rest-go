package middlewares

import (
	"fmt"

	structs "github.com/Jason-Bai/open-rest-go/structs/middlewares"
	"github.com/gin-gonic/gin"
)

// Context reset context for every request
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		rc := structs.NewRequestContext(c)

		c.Set("Context", rc)

		fmt.Println("before context middleware")

		c.Next()

		fmt.Println("after context middleware")
	}
}
