package model

import "time"

type Task struct {
	ID              uint   `gorm:"primarykey"`
	ProjectId       string `gorm:"index"`
	TaskTitle       string
	TaskDescription string
	TaskType        string
	StartDate       string    `gorm:"type:date"`
	EndDate         string    `gorm:"type:date"`
	CreatedDate     time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime:milli"`
	ProjectFkey     Project   `gorm:"foreignKey:ProjectId"`
}

func (Task) TableName() string {
	return "tb_task"
}
