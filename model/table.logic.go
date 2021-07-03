package model

import "time"

type Logic struct {
	ID           uint   `gorm:"primarykey"`
	Operation    string `gorm:"type:varchar(2)"`
	QuestionId   uint   `gorm:"index"`
	CompareValue string
	UpdatedDate  time.Time `gorm:"autoUpdateTime:milli"`
	CreatedDate  time.Time `gorm:"autoCreateTime:milli"`

	QuestionFkey Question `gorm:"foreignKey:QuestionId"`
}

func (Logic) TableName() string {
	return "tb_logic"
}
