package model

import (
	"errors"

	"github.com/fatih/structs"
)

type ProjectForm struct {
	ID                 string `json:"id" form:"id"`
	StartDate          string `json:"startDate" form:"startDate"`
	EndDate            string `json:"endDate" form:"endDate"`
	ProjectTitle       string `json:"projectTitle" form:"projectTitle"`
	RegisterId         uint   `json:"registerId" form:"registerId"`
	ProjectDescription string `json:"projectDescription" form:"projectDescription"`
}

type PagingProjectForm struct {
	TotalCount  int64          `json:"totalCount" form:"totalCount"`
	ProjectForm *[]ProjectForm `json:"projectList" form:"projectList"`
}

func (ProjectForm) getTableName() string {
	return "tb_project"
}

type ProjectQuerySet struct {
}

func (ProjectQuerySet) getTableName() string {
	return "tb_project"
}

func (p ProjectQuerySet) SelectOne(id interface{}) (res *ProjectForm, err error) {
	project := &ProjectForm{}

	result := Gdb.Table(p.getTableName()).Where("id = ?", id.(string)).Find(&project)

	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (p ProjectQuerySet) Select(param interface{}) (res *PagingProjectForm, err error) {

	var pt *Pagetype
	var pok, ucok bool
	var pf *ProjectForm
	var totalCount int64

	project := &[]ProjectForm{}
	s := structs.New(param)
	resForm := &PagingProjectForm{}

	PageInfo, pfok := s.FieldOk("PageInfo")
	ProjectFilter, ufok := s.FieldOk("ProjectFilter")

	if !pfok || PageInfo == nil {
		pt = defaultPageInfo
	} else {
		pt, pok = PageInfo.Value().(*Pagetype)
		if !pok || pt == nil {
			pt = defaultPageInfo
		}
	}

	if !ufok {
		pf = &ProjectForm{}
	} else {
		pf, ucok = ProjectFilter.Value().(*ProjectForm)
		if !ucok {
			pf = &ProjectForm{}
		}
	}

	result := Gdb.Scopes(Paginate(pt)).Table(p.getTableName()).Where(pf).Find(&project)

	if result.Error != nil {
		return nil, result.Error
	}

	Gdb.Table(p.getTableName()).Where(pf).Count(&totalCount)

	resForm.TotalCount = totalCount
	resForm.ProjectForm = project

	return resForm, nil
}

func (p ProjectQuerySet) InsertOne(data *Project) error {
	var count int64
	Gdb.Table(p.getTableName()).Where("project_title = ?", data.ProjectTitle).Count(&count)

	if count > 0 {
		return errors.New("already exist")
	}

	result := Gdb.Create(data)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (p ProjectQuerySet) InsertMany(data *[]Project) error {
	result := Gdb.CreateInBatches(data, len(*data))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p ProjectQuerySet) DeleteOne(id interface{}) error {
	dbresult := Gdb.Where("id = ?", id.(string)).Delete(&Project{})

	if dbresult.Error != nil {
		return dbresult.Error
	}

	if dbresult.RowsAffected == 0 {
		return errors.New("can't find the user")
	}

	return nil
}

func (p ProjectQuerySet) UpdateOne(id string, param *ProjectForm) error {

	dbresult := Gdb.Table(p.getTableName()).Where("id = ?", id).Updates(*param)

	if dbresult.Error != nil {
		return dbresult.Error
	}

	if dbresult.RowsAffected == 0 {
		return errors.New("can't find the user")
	}

	return nil
}
