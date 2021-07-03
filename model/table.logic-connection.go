package model

import "time"

type LogicConnection struct {
	ID             uint `gorm:"primarykey"`
	IsPageLogic    bool
	FirstLogicId   uint
	LastLogicId    uint
	LogicOperation string `gorm:"type:varchar(2)"`

	CreatedDate time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate time.Time `gorm:"autoUpdateTime:milli"`

	LogicFkey  Logic `gorm:"foreignKey:FirstLogicId"`
	LogicFkey2 Logic `gorm:"foreignKey:LastLogicId"`
}

func (LogicConnection) TableName() string {
	return "tb_logic_connection"
}
