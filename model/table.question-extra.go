package model

import "time"

type QuestionExtra struct {
	ID           uint `gorm:"primarykey"`
	QuestionId   uint `gorm:"index"`
	StaticList   string
	IsStatic     bool
	HardCodingJs string
	IsHardCoding bool
	CreatedDate  time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime:milli"`

	QuestionFkey Question `gorm:"foreignKey:QuestionId"`
}

func (QuestionExtra) TableName() string {
	return "tb_question_extra"
}
