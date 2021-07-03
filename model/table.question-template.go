package model

import "time"

type QuestionTemplate struct {
	ID                 uint      `gorm:"primarykey"`
	QuestionType       string    `gorm:"type:varchar(1)"`
	TitleFcolor        string    `gorm:"type:varchar(20)"`
	TitleBgcolor       string    `gorm:"type:varchar(20)"`
	DescriptionFcolor  string    `gorm:"type:varchar(20)"`
	DescriptionBgcolor string    `gorm:"type:varchar(20)"`
	LayoutId           uint      `gorm:"index"`
	ContentId          uint      `gorm:"index"`
	CreatedDate        time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime:milli"`

	QuestionLayoutFkey  QuestionLayout  `gorm:"foreignKey:LayoutId"`
	QuestionContentFkey QuestionContent `gorm:"foreignKey:ContentId"`
	QuestionTypeFkey    QuestionType    `gorm:"foreignKey:QuestionType"`
}

func (QuestionTemplate) TableName() string {
	return "tb_question_template"
}
