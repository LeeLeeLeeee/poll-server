package model

import "time"

type Project struct {
	ID                 string `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	StartDate          string `gorm:"type:date"`
	EndDate            string `gorm:"type:date"`
	ProjectTitle       string `gorm:"type:varchar(250)"`
	RegisterId         uint
	ProjectDescription string
	ProjectStatus      string    `gorm:"type:varchar(1);check:project_status_checker, project_status in ('1', '2', '3'); default:'1' "`
	CreatedDate        time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime:milli"`
	User               User      `gorm:"foreignKey:RegisterId"`
}

func (Project) TableName() string {
	return "tb_project"
}
