package http

import (
	"github.com/Jason-Bai/open-rest-go/apis/user"
	"github.com/Jason-Bai/open-rest-go/apis/user/repository"
	"github.com/Jason-Bai/open-rest-go/apis/user/usecase"
	"github.com/Jason-Bai/open-rest-go/helper/getter"
	"github.com/gin-gonic/gin"
)

type (
	userResource struct {
		usecase user.UseCase
	}
)

// NewUserHttpRoutes add user http routes
func NewUserHttpRoutes(c *gin.Engine) {
	ur := repository.NewRepository()

	usecase := usecase.NewUseCase(ur)

	r := &userResource{
		usecase,
	}

	c.POST("/users", []gin.HandlerFunc{
		getter.Getter(nil, ""),
		r.create,
	}...)
}

func (ur *userResource) create(c *gin.Context) {
	data := [3]int{1, 2, 3}
	c.JSON(200, gin.H{
		"total": len(data),
		"data":  data,
	})
}
