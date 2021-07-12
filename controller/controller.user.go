package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/lib"
	"github.com/leeleeleeee/web-app/model"
)

type UserController struct {
	UserQuerySet model.UserQuerySet
}

func (uc UserController) Init(v1 *gin.RouterGroup, e ...interface{}) {
	v1.GET("/user", uc.getUser)
	v1.GET("/user/:id", uc.getUserOne)
	v1.POST("/user", uc.insertUser)
	v1.POST("/userMany", uc.insertUserMany)
	v1.DELETE("/user/:id", uc.deleteOne)
	v1.PUT("/user/:id", uc.updateOne)

}

func (u UserController) getUser(c *gin.Context) {
	type userParam struct {
		PageInfo    *model.Pagetype
		UserFileter *model.UserForm
	}

	var param *userParam

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := u.UserQuerySet
	result, err := qs.Select(param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)

}

func (u UserController) getUserOne(c *gin.Context) {
	id := c.Param("id")
	qs := u.UserQuerySet
	result, err := qs.SelectOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u UserController) insertUser(c *gin.Context) {
	var json *model.User
	var encryptErr error

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.Password, encryptErr = lib.EncryptSha256(json.Password)

	if encryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": encryptErr.Error()})
		return
	}

	qs := u.UserQuerySet
	err := qs.InsertOne(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (u UserController) insertUserMany(c *gin.Context) {
	var json *[]model.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := u.UserQuerySet
	err := qs.InsertMany(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (u UserController) deleteOne(c *gin.Context) {
	id := c.Param("id")
	qs := u.UserQuerySet
	err := qs.DeleteOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "success")
}

func (u UserController) updateOne(c *gin.Context) {
	var json *model.UserForm
	id := c.Param("id")
	qs := u.UserQuerySet

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := qs.UpdateOne(id, json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}
