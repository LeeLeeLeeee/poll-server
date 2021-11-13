package model

import "time"

type QuestionContent struct {
	ID                 uint `gorm:"primarykey"`
	QuestionTemplateId uint
	IsLabel            bool
	LabelPlacement     string `gorm:"type:varchar(5)"`
	LabelSize          string `gorm:"type:varchar(6)"`
	IsValue            bool
	ValuePlacement     string `gorm:"type:varchar(5)"`
	ValueSize          string `gorm:"type:varchar(6)"`
	IsMultiColumn      bool
	MultiColumnNum     string `gorm:"type:varchar(2)"`
	IsValueShuffle     bool
	ValueShuffleFix    int `gorm:"type:smallint"`
	IsAttrShuffle      bool
	AttrShuffleFix     int `gorm:"type:smallint"`
	AttrViewInOrder    bool
	IsAttrLabelMerge   bool
	IsAttrCodeHeader   bool
	IsDivideAttrLabel  bool
	CreatedDate        time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate        time.Time `gorm:"autoUpdateTime:milli"`

	QuestionTemplateFkey QuestionTemplate `gorm:"foreignKey:QuestionTemplateId"`
}

func (QuestionContent) TableName() string {
	return "tb_question_content"
}
