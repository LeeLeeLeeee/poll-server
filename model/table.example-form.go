package model

import "time"

type ExampleForm struct {
	ID         uint   `gorm:"primarykey"`
	QuestionId uint   `gorm:"index"`
	AttrId     uint   `gorm:"index"`
	FormLabel  string `gorm:"type:varchar(250)"`
	FormCode   string `gorm:"type:varchar(5)"`
	FormValue  string `gorm:"index"`
	IsOpen     bool
	IsEtc      bool
	IsNone     bool

	Question     Question  `gorm:"foreignKey:QuestionId"`
	CreatedDate  time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime:milli"`
	FormAttrFkey FormAttr  `gorm:"foreignKey:AttrId"`
	QuestionFkey Question  `gorm:"foreignKey:QuestionId"`
}

func (ExampleForm) TableName() string {
	return "tb_example_form"
}
