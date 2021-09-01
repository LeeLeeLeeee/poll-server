package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (ac AuthController) Init(v1 *gin.RouterGroup, e ...interface{}) {
	v1.POST("/login", LogIn)
	v1.POST("/logout", LogOut)
	v1.GET("/checkToken", CheckJwt())
}
