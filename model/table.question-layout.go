package model

import "time"

type QuestionLayout struct {
	ID                 uint `gorm:"primarykey"`
	QuestionTemplateId uint
	IsHeader           bool
	IsBottom           bool
	LayoutPlacement    string `gorm:"type:varchar(1)"`
	HeaderFcolor       string `gorm:"type:varchar(20)"`
	HeaderBgcolor      string `gorm:"type:varchar(20)"`
	HeaderLogoUrl      string
	HeaderText         string
	HeaderAlign        string `gorm:"type:varchar(5)"`
	BottomFcolor       string `gorm:"type:varchar(20)"`
	BottomBgcolor      string `gorm:"type:varchar(20)"`
	BottomLogoUrl      string
	BottomText         string
	BottomAlign        string    `gorm:"type:varchar(5)"`
	CreatedDate        time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime:milli"`

	QuestionTemplateFkey QuestionTemplate `gorm:"foreignKey:QuestionTemplateId"`
}

func (QuestionLayout) TableName() string {
	return "tb_question_layout"
}
