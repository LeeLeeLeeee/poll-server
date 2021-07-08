package model

import (
	"time"
)

type Question struct {
	ID                  uint   `gorm:"primarykey"`
	UserId              uint   `gorm:"index"`
	ProjectId           string `gorm:"type:uuid"`
	TaskId              uint   `gorm:"index"`
	QuestionCode        string `gorm:"index"`
	QuestionType        string `gorm:"type:smallint"`
	QuestionTemplateId  uint
	QuestionTitle       string
	QuestionDescription string
	IsActive            bool
	IsEntryLogic        bool
	IsExitLogic         bool
	CreatedDate         time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate         time.Time `gorm:"autoUpdateTime:milli"`

	UserFkey    User    `gorm:"foreignKey:UserId"`
	ProjectFkey Project `gorm:"foreignKey:ProjectId"`
	TaskFkey    Task    `gorm:"foreignKey:TaskId"`
}

func (Question) TableName() string {
	return "tb_question"
}
