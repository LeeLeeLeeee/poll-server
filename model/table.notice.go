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

	CreatedDate   time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime:milli"`
	ToUsersFkey   User      `gorm:"foreignKey:ToUserId"`
	FromUsersFkey User      `gorm:"foreignKey:FromUserId"`
	ProjectFkey   Project   `gorm:"foreignKey:ProjectId"`
}

func (Notice) TableName() string {
	return "tb_notice"
}
