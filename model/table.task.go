package model

import "time"

type Task struct {
	ID              uint `gorm:"primarykey"`
	ProjectID       uint `gorm:"index"`
	TaskTitle       string
	TaskDescription string
	TaskType        string
	StartDate       string    `gorm:"type:date"`
	EndDate         string    `gorm:"type:date"`
	CreatedDate     time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime:milli"`
}

func (Task) TableName() string {
	return "tb_task"
}
