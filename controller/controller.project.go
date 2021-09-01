package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/model"
)

type ProjectController struct {
	ProjectQuerySet model.ProjectQuerySet
}

func (c ProjectController) Init(v1 *gin.RouterGroup, e ...interface{}) {

	v1.GET("/project", c.getProject)
	v1.GET("/project/:id", c.getProjectOne)
	v1.POST("/project", c.insertProject)
	v1.POST("/projectMany", c.insertProjectMany)
	v1.DELETE("/project/:id", c.deleteOne)
	v1.PUT("/project/:id", c.updateOne)

}

func (p ProjectController) getProject(c *gin.Context) {
	type projectParam struct {
		PageInfo      *model.Pagetype
		ProjectFilter *model.ProjectForm
	}

	var param *projectParam

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := p.ProjectQuerySet
	result, err := qs.Select(param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)

}

func (u ProjectController) getProjectOne(c *gin.Context) {
	id := c.Param("id")
	qs := u.ProjectQuerySet
	result, err := qs.SelectOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u ProjectController) insertProject(c *gin.Context) {
	var json *model.Project

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := u.ProjectQuerySet
	err := qs.InsertOne(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (u ProjectController) insertProjectMany(c *gin.Context) {
	var json *[]model.Project

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := u.ProjectQuerySet
	err := qs.InsertMany(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (u ProjectController) deleteOne(c *gin.Context) {
	id := c.Param("id")
	qs := u.ProjectQuerySet
	err := qs.DeleteOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "success")
}

func (u ProjectController) updateOne(c *gin.Context) {
	var json *model.ProjectForm
	id := c.Param("id")
	qs := u.ProjectQuerySet

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
