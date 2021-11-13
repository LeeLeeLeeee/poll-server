package model

type QuestionType struct {
	ID       uint   `gorm:"primarykey; type:smallint"`
	TypeName string `gorm:"type:varchar(20)"`
}

func (QuestionType) TableName() string {
	return "tb_question_type"
}
