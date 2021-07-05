package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc UserController) Init(v1 *gin.RouterGroup, e ...interface{}) {
	v1.GET("/someGet", uc.someGetFunc)
	v1.GET("/someGet2", uc.someGetFunc)
	v1.GET("/someGet3", uc.someGetFunc)
}

func (UserController) someGetFunc(c *gin.Context) {
	example := c.MustGet("userID").(uint64)
	responseContent := "Hello " + strconv.Itoa(int(example))
	c.String(http.StatusOK, responseContent)
}
