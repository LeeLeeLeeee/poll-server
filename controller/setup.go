package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Init(router *gin.RouterGroup, e ...interface{})
}

func DoInit(rg *gin.RouterGroup, c Controller, e ...interface{}) {
	c.Init(rg, e...)
}
