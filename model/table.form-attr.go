package model

import (
	"time"
)

type FormAttr struct {
	ID             uint `gorm:"primarykey"`
	QuestionId     uint `gorm:"index"`
	ExampleId      uint `gorm:"index"`
	AttrLabel      string
	AttrLevel      uint `gorm:"type:smallint"`
	AttrCode       uint `gorm:"type:smallint"`
	IsValueReverse bool
	CreatedDate    time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate    time.Time `gorm:"autoUpdateTime:milli"`

	QuestionFKey Question    `gorm:"foreignKey:QuestionId"`
	ExampleKey   ExampleForm `gorm:"foreignKey:ExampleId"`
}

func (FormAttr) TableName() string {
	return "tb_form_attr"
}
