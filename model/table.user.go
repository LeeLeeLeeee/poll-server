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
	IsStaff     string    `gorm:"<-"`
	IsActive    string    `gorm:"<-"`
	LastLogin   time.Time `gorm:"<-"`
	IsSuperuser bool
}

func (User) TableName() string {
	return "tb_user"
}
