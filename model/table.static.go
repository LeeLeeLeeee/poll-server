package model

import "time"

type Static struct {
	ID           uint   `gorm:"primarykey"`
	ProjectId    string `gorm:"index;type:uuid"`
	TaskId       uint   `gorm:"index"`
	StaticType   string
	CreatedDate  time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime:milli"`
	TaskFkey     Task      `gorm:"foreignKey:TaskId"`
	ProejectFkey Project   `gorm:"foreignKey:ProjectId"`
}

func (Static) TableName() string {
	return "tb_static"
}
