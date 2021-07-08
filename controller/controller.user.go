package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/model"
)

type UserController struct {
	UserQuerySet model.UserQuerySet
}

func (uc UserController) Init(v1 *gin.RouterGroup, e ...interface{}) {
	v1.GET("/user/:id", uc.getUserId)
}

// func (UserController) someGetFunc(c *gin.Context) {
// 	example := c.MustGet("userID").(uint64)
// 	responseContent := "Hello " + strconv.Itoa(int(example))
// 	c.String(http.StatusOK, responseContent)
// }

func (u UserController) getUserId(c *gin.Context) {
	id := c.Param("id")
	qs := u.UserQuerySet
	qs.SelectOne(id)
}
