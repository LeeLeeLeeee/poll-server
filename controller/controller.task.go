package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeleeleeee/web-app/model"
)

type TaskController struct {
	TaskQuerySet model.TaskQuerySet
}

func (c TaskController) Init(v1 *gin.RouterGroup, e ...interface{}) {

	v1.GET("/task", c.getTask)
	v1.GET("/task/:id", c.getTaskOne)
	v1.POST("/task", c.insertTask)
	v1.POST("/taskMany", c.insertTaskMany)
	v1.DELETE("/task/:id", c.deleteOne)
	v1.PUT("/task/:id", c.updateOne)

}

func (p TaskController) getTask(c *gin.Context) {
	type taskParam struct {
		PageInfo   *model.Pagetype
		TaskFilter *model.TaskForm
	}

	var param *taskParam

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := p.TaskQuerySet
	result, err := qs.Select(param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)

}

func (u TaskController) getTaskOne(c *gin.Context) {
	id := c.Param("id")
	qs := u.TaskQuerySet
	result, err := qs.SelectOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (u TaskController) insertTask(c *gin.Context) {
	var json *model.Task

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := u.TaskQuerySet
	err := qs.InsertOne(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (u TaskController) insertTaskMany(c *gin.Context) {
	var json *[]model.Task

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs := u.TaskQuerySet
	err := qs.InsertMany(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (u TaskController) deleteOne(c *gin.Context) {
	id := c.Param("id")
	qs := u.TaskQuerySet
	err := qs.DeleteOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "success")
}

func (u TaskController) updateOne(c *gin.Context) {
	var json *model.TaskForm
	id := c.Param("id")
	qs := u.TaskQuerySet

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
