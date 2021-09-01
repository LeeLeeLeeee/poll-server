package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UserId      string    `gorm:"not_null;uniqueindex" json:"user_id" `
	Password    string    `gorm:"<-" json:"password" `
	FirstName   string    `gorm:"<-" json:"firstname" `
	LastName    string    `gorm:"<-" json:"lastname" `
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	LastLogin   time.Time `gorm:"autoCreateTime" json:"last_login"`
	CreatedDate time.Time `gorm:"autoUpdateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoCreateTime" json:"update_date"`
	LoginCount  uint      `gorm:"type:smallint; default:0" json:"login_count"`
	IsSuperuser bool      `json:"is_superuser"`
}

func (User) TableName() string {
	return "tb_user"
}

func (u User) BeforeCreate(tx *gorm.DB) error {
	/*
		Something Logic..
	*/
	return nil
}
