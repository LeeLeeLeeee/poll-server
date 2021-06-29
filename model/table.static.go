package model

import "time"

type Static struct {
	ID          uint `gorm:"primarykey"`
	ProjectId   uint `gorm:"index"`
	TaskId      uint `gorm:"index"`
	FileType    string
	FileUrl     string
	CreatedDate time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate time.Time `gorm:"autoUpdateTime:milli"`
}

func (Static) TableName() string {
	return "tb_static"
}
