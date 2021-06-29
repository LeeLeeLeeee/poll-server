package model

type Project struct {
	ID                 uint   `gorm:"primarykey"`
	ProjectId          uint   `gorm:"index"`
	StartDate          string `gorm:"type:date"`
	EndDate            string `gorm:"type:date"`
	ProjectTitle       string
	RegisterId         uint
	CreatedDate        string `gorm:"autoCreateTime:milli"`
	ProjectType        string
	ProjectDescription string
	ProjectStatus      string
}

func (Project) TableName() string {
	return "tb_project"
}
