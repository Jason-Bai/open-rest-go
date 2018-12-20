package structs

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// RequestContext define a new request context
type RequestContext interface {
	RequestID() string
	SetTx(tx *gorm.DB)
	Tx() *gorm.DB
	RC() *gin.Context
}

type requestContext struct {
	requestID string
	tx        *gorm.DB
	rc        *gin.Context
}

// UserId implemented requestContext's RequestContext interface UserID method
func (rc requestContext) RequestID() string {
	return rc.requestID
}

func (rc requestContext) SetTx(tx *gorm.DB) {
	rc.tx = tx
}

func (rc requestContext) Tx() *gorm.DB {
	return rc.tx
}

// RC implemented requestContext's RequestContext interface RC method
func (rc requestContext) RC() *gin.Context {
	return rc.rc
}

// NewRequestContext create a new request context for every request
func NewRequestContext(c *gin.Context) RequestContext {
	requestID := c.Request.Header.Get("X-Request-ID")

	return &requestContext{
		requestID: requestID,
		rc:        c,
	}
}

// GetRequestContext get a request context
func GetRequestContext(c *gin.Context) (RequestContext, error) {
	ctx, exists := c.Get("Context")

	if exists == false {
		return nil, fmt.Errorf("no request context")
	}

	rc := ctx.(RequestContext)

	return rc, nil
}
