package model

import (
	"time"

	"gorm.io/datatypes"
)

type Question struct {
	ID           uint   `gorm:"primarykey"`
	UserId       uint   `gorm:"index"`
	ProjectId    uint   `gorm:"index"`
	QuestionCode string `gorm:"index"`
	QuestionType string
	IsActive     bool
	LogicBefore  datatypes.JSON
	LogicAfter   datatypes.JSON
	CreatedDate  time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime:milli"`
}

func (Question) TableName() string {
	return "tb_question"
}
