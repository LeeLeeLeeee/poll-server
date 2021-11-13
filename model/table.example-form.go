package model

import "time"

type ExampleForm struct {
	ID         uint   `gorm:"primarykey"`
	QuestionId uint   `gorm:"index"`
	FormLabel  string `gorm:"type:varchar(250)"`
	FormCode   string `gorm:"type:varchar(5)"`
	FormValue  string `gorm:"index"`
	IsOpen     bool
	IsEtc      bool
	IsNone     bool

	CreatedDate  time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime:milli"`
	QuestionFkey Question  `gorm:"foreignKey:QuestionId"`
}

func (ExampleForm) TableName() string {
	return "tb_example_form"
}
