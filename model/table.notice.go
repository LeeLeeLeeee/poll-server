package model

import (
	"time"
)

type Notice struct {
	ID         uint `gorm:"primarykey"`
	ToUserId   uint `gorm:"index"`
	FromUserId uint `gorm:"index"`
	ProjectId  uint `gorm:"index"`
	TaskId     uint `gorm:"index"`
	QuestionId uint `gorm:"index"`
	IsError    bool
	IsChecked  bool
	UpdateDate time.Time `gorm:"autoUpdateTime:milli"`
}

func (Notice) TableName() string {
	return "tb_notice"
}
