package model

import (
	"time"

	"gorm.io/datatypes"
)

type Data struct {
	ID          uint `gorm:"primarykey"`
	UserId      uint `gorm:"index"`
	ProjectId   uint `gorm:"index"`
	TaskId      uint `gorm:"index"`
	QuestionId  uint `gorm:"index"`
	SurveyData  datatypes.JSON
	DateHistory datatypes.JSON
	CreatedDate time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate time.Time `gorm:"autoUpdateTime:milli"`
}

func (Data) TableName() string {
	return "tb_data"
}
