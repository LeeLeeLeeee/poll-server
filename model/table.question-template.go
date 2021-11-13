package model

import "time"

type QuestionTemplate struct {
	ID                 uint      `gorm:"primarykey"`
	QuestionId         uint      `gorm:"index"`
	TitleFcolor        string    `gorm:"type:varchar(20)"`
	TitleBgcolor       string    `gorm:"type:varchar(20)"`
	DescriptionFcolor  string    `gorm:"type:varchar(20)"`
	DescriptionBgcolor string    `gorm:"type:varchar(20)"`
	CreatedDate        time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime:milli"`

	QuestionFkey Question `gorm:"foreignKey:QuestionId"`
}

func (QuestionTemplate) TableName() string {
	return "tb_question_template"
}
