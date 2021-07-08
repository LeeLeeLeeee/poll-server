package model

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primarykey"`
	UserId      string    `gorm:"not_null;uniqueindex"`
	Password    string    `gorm:"<-"`
	Email       string    `gorm:"<-"`
	FirstName   string    `gorm:"<-"`
	LastName    string    `gorm:"<-"`
	IsActive    bool      `gorm:"default:true"`
	LastLogin   time.Time `gorm:"<-"`
	CreatedDate time.Time `gorm:"autoCreateTime:milli"`
	UpdatedDate time.Time `gorm:"autoUpdateTime:milli"`
	LoginCount  uint      `gorm:"type:smallint; default:0"`
	IsSuperuser bool
}

func (User) TableName() string {
	return "tb_user"
}
