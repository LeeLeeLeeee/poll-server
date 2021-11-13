package model

import (
	"errors"

	"github.com/fatih/structs"
)

type TaskForm struct {
	ID         string `json:"id" form:"id"`
	ProjectId  string `json:"projectId" form:"projectId"`
	StartDate  string `json:"startDate" form:"startDate"`
	EndDate    string `json:"endDate" form:"endDate"`
	TaskTitle  string `json:"taskTitle" form:"taskTitle"`
	TaskStatus string `json:"taskStatus" form:"taskStatus"`
	TaskType   string `json:"taskType" form:"taskType"`
	RegisterId uint   `json:"registerId" form:"registerId"`
}

type PagingTaskForm struct {
	TotalCount int64       `json:"totalCount" form:"totalCount"`
	TaskForm   *[]TaskForm `json:"taskList" form:"taskList"`
}

type TaskQuerySet struct {
}

func (TaskQuerySet) getTableName() string {
	return "tb_task"
}

func (p TaskQuerySet) SelectOne(id interface{}) (res *TaskForm, err error) {
	task := &TaskForm{}

	result := Gdb.Table(p.getTableName()).Where("id = ?", id.(string)).Find(&task)

	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (p TaskQuerySet) Select(param interface{}) (res *PagingTaskForm, err error) {

	var pt *Pagetype
	var pok, ucok bool
	var pf *TaskForm
	var totalCount int64

	task := &[]TaskForm{}
	s := structs.New(param)
	resForm := &PagingTaskForm{}

	PageInfo, pfok := s.FieldOk("PageInfo")
	taskFilter, ufok := s.FieldOk("TaskFilter")

	if !pfok || PageInfo == nil {
		pt = defaultPageInfo
	} else {
		pt, pok = PageInfo.Value().(*Pagetype)
		if !pok || pt == nil {
			pt = defaultPageInfo
		}
	}

	if !ufok {
		pf = &TaskForm{}
	} else {
		pf, ucok = taskFilter.Value().(*TaskForm)
		if !ucok {
			pf = &TaskForm{}
		}
	}

	result := Gdb.Scopes(Paginate(pt)).Table(p.getTableName()).Where(pf).Order("created_date desc").Find(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	Gdb.Table(p.getTableName()).Where(pf).Count(&totalCount)

	resForm.TotalCount = totalCount
	resForm.TaskForm = task

	return resForm, nil
}

func (p TaskQuerySet) InsertOne(data *Task) error {
	var count int64
	Gdb.Table(p.getTableName()).Where("task_title = ?", data.TaskTitle).Count(&count)

	if count > 0 {
		return errors.New("already exist")
	}

	result := Gdb.Create(data)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (p TaskQuerySet) InsertMany(data *[]Task) error {
	result := Gdb.CreateInBatches(data, len(*data))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p TaskQuerySet) DeleteOne(id interface{}) error {
	dbresult := Gdb.Where("id = ?", id.(string)).Delete(&Task{})

	if dbresult.Error != nil {
		return dbresult.Error
	}

	if dbresult.RowsAffected == 0 {
		return errors.New("can't find the user")
	}

	return nil
}

func (p TaskQuerySet) UpdateOne(id string, param *TaskForm) error {

	dbresult := Gdb.Table(p.getTableName()).Where("id = ?", id).Updates(*param)

	if dbresult.Error != nil {
		return dbresult.Error
	}

	if dbresult.RowsAffected == 0 {
		return errors.New("can't find the user")
	}

	return nil
}
