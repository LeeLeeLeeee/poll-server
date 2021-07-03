package model

import "time"

type ExampleLogic struct {
	ID              uint   `gorm:"primarykey"`
	LoadType        string `gorm:"type:varchar(1)"`
	LogicId         uint   `gorm:"index"`
	BaseExampleId   uint
	BaseAttrId      uint
	TargetExampleId uint
	TargetAttrId    uint
	ToHide          bool
	ToShow          bool
	ToDisabled      bool
	ToEnabled       bool
	CreatedDate     time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate     time.Time `gorm:"autoUpdateTime:milli"`

	ExampleFormFkey  ExampleForm `gorm:"foreignKey:BaseExampleId"`
	ExampleFormFkey2 ExampleForm `gorm:"foreignKey:TargetExampleId"`
	FormAttrIdFkey   FormAttr    `gorm:"foreignKey:BaseAttrId"`
	FormAttrIdFkey2  FormAttr    `gorm:"foreignKey:TargetAttrId"`
	LogicFkey        Logic       `gorm:"foreignKey:LogicId"`
}

func (ExampleLogic) TableName() string {
	return "tb_example_logic"
}
