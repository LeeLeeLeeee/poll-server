package model

import "time"

type QuestionLogic struct {
	ID          uint   `gorm:"primarykey"`
	LoadType    string `gorm:"type:varchar(1)"`
	LogicId     uint   `gorm:"index"`
	LogicWeight uint   `gorm:"type:smallint"`
	ExampleId   uint   `gorm:"index"`
	AttrId      uint   `gorm:"index"`
	IsEntry     bool
	UpdatedDate time.Time `gorm:"autoUpdateTime:milli"`
	CreatedDate time.Time `gorm:"autoCreateTime:milli"`

	LogicFkey   Logic       `gorm:"foreignKey:LogicId"`
	AttrFkey    FormAttr    `gorm:"foreignKey:AttrId"`
	ExampleFkey ExampleForm `gorm:"foreignKey:ExampleId"`
}

func (QuestionLogic) TableName() string {
	return "tb_question_logic"
}
