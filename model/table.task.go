package model

import "time"

type Task struct {
	ID          uint   `gorm:"primarykey"`
	ProjectId   string `gorm:"index"`
	TaskTitle   string
	TaskType    string
	TaskStatus  string `gorm:"type:varchar(1);check:task_status_checker, task_status in ('1', '2', '3'); default:'1' "`
	RegisterId  uint
	StartDate   string    `gorm:"type:date"`
	EndDate     string    `gorm:"type:date"`
	CreatedDate time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate time.Time `gorm:"autoUpdateTime:milli"`
	ProjectFkey Project   `gorm:"foreignKey:ProjectId"`
	User        User      `gorm:"foreignKey:RegisterId"`
}

func (Task) TableName() string {
	return "tb_task"
}
